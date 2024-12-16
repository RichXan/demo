package impl

import (
	"context"
	"errors"
	"fmt"
	"time"

	"x-micro-blog/internal/access/model"
	"x-micro-blog/internal/micro/service"

	"github.com/richxan/xcommon/xauth"
	"github.com/richxan/xcommon/xerror"
	"github.com/richxan/xcommon/xlog"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token expired")
)

type authService struct {
	db          *gorm.DB
	logger      *xlog.Logger
	redisClient *redis.Client
	oauthConfig *xauth.OAuthConfig
	tokenStore  xauth.TokenStore
}

// NewAuthService 创建新的认证服务
func NewAuthService(db *gorm.DB, logger *xlog.Logger, redisClient *redis.Client, oauthConfig *xauth.OAuthConfig) service.AuthService {
	return &authService{
		db:          db,
		logger:      logger,
		redisClient: redisClient,
		oauthConfig: oauthConfig,
		tokenStore:  xauth.NewRedisTokenStore(redisClient),
	}
}

// RefreshToken 刷新令牌
func (s *authService) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	// 解析刷新令牌
	claims, err := xauth.ParseRefreshToken(refreshToken)
	if err != nil {
		if errors.Is(err, xauth.ErrExpiredToken) {
			return "", xerror.TokenExpired
		}
		return "", xerror.RefreshTokenInvalid
	}

	// 检查令牌是否被撤销
	if s.tokenStore.IsTokenRevoked(ctx, claims.TokenID) {
		return "", xerror.RefreshTokenInvalid
	}

	// 生成新的令牌对
	tokenPair, err := xauth.GenerateTokenPair(claims.UserID, claims.Username)
	if err != nil {
		return "", xerror.Wrap(err, xerror.CodeSystemError, "failed to generate token pair")
	}

	// 撤销旧的刷新令牌
	if err := s.tokenStore.RevokeToken(ctx, claims.TokenID, time.Hour*24*7); err != nil {
		s.logger.Error().Err(err).Msg("failed to revoke old refresh token")
	}

	return tokenPair.AccessToken, nil
}

// RevokeToken 撤销令牌
func (s *authService) RevokeToken(ctx context.Context, token string) error {
	// 解析令牌
	claims, err := xauth.ParseAccessToken(token)
	if err != nil {
		return xerror.TokenInvalid
	}

	// 将令牌加入黑名单
	if err := s.tokenStore.RevokeToken(ctx, claims.TokenID, time.Hour*24); err != nil {
		return xerror.Wrap(err, xerror.CodeSystemError, "failed to revoke token")
	}

	return nil
}

// ValidateToken 验证令牌
func (s *authService) ValidateToken(ctx context.Context, token string) (int64, error) {
	// 解析令牌
	claims, err := xauth.ParseAccessToken(token)
	if err != nil {
		if errors.Is(err, xauth.ErrExpiredToken) {
			return 0, xerror.TokenExpired
		}
		return 0, xerror.TokenInvalid
	}

	// 检查令牌是否被撤销
	if s.tokenStore.IsTokenRevoked(ctx, claims.TokenID) {
		return 0, xerror.TokenInvalid
	}

	// 检查用户是否存在
	var user model.User
	if err := s.db.First(&user, claims.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, xerror.UserNotFound
		}
		return 0, xerror.Wrap(err, xerror.CodeSystemError, "failed to get user")
	}

	return int64(claims.UserID), nil
}

// GetOAuthRedirectURL 获取OAuth重定向URL
func (s *authService) GetOAuthRedirectURL(ctx context.Context, provider string) (string, error) {
	// 获取OAuth提供商
	p, err := s.oauthConfig.GetProvider(provider)
	if err != nil {
		return "", xerror.OAuthFailed
	}

	// 生成状态值
	state := xauth.GenerateState()

	// 生成授权URL
	url := p.Config.AuthCodeURL(state)

	return url, nil
}

// HandleOAuthCallback 处理OAuth回调
func (s *authService) HandleOAuthCallback(ctx context.Context, provider, code, state string) (string, error) {
	// 获取OAuth提供商
	p, err := s.oauthConfig.GetProvider(provider)
	if err != nil {
		return "", xerror.OAuthFailed
	}

	// 交换授权码获取令牌
	token, err := p.Config.Exchange(ctx, code)
	if err != nil {
		return "", xerror.Wrap(err, xerror.CodeSystemError, "failed to exchange OAuth code")
	}

	// 获取用户信息
	client := p.Config.Client(ctx, token)
	userInfo, err := p.GetUserInfo(ctx, client)
	if err != nil {
		return "", xerror.Wrap(err, xerror.CodeSystemError, "failed to get OAuth user info")
	}

	// 查找或创建社交账号
	var socialAccount model.SocialAccount
	err = s.db.Where("provider = ? AND open_id = ?", provider, userInfo.OpenID).First(&socialAccount).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", xerror.Wrap(err, xerror.CodeSystemError, "failed to query social account")
		}

		// 创建新用户
		user := &model.User{
			Username: fmt.Sprintf("%s_%s", provider, userInfo.ID),
			Nickname: userInfo.Name,
			Avatar:   userInfo.AvatarURL,
			Email:    userInfo.Email,
		}

		// 创建社交账号
		socialAccount = model.SocialAccount{
			Provider: provider,
			OpenID:   userInfo.OpenID,
			UnionID:  userInfo.UnionID,
			Nickname: userInfo.Name,
			Avatar:   userInfo.AvatarURL,
			Extra:    userInfo.Extra,
			User:     user,
		}

		if err := s.db.Create(&socialAccount).Error; err != nil {
			return "", xerror.Wrap(err, xerror.CodeSystemError, "failed to create social account")
		}
	}

	// 生成令牌对
	tokenPair, err := xauth.GenerateTokenPair(uint64(socialAccount.UserID), socialAccount.User.Username)
	if err != nil {
		return "", xerror.Wrap(err, xerror.CodeSystemError, "failed to generate token pair")
	}

	return tokenPair.AccessToken, nil
}
