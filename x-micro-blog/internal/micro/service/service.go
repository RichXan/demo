package service

import (
	"context"
)

// AuthService 认证服务接口
type AuthService interface {
	// Token管理
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
	RevokeToken(ctx context.Context, token string) error
	ValidateToken(ctx context.Context, token string) (int64, error)

	// OAuth管理
	GetOAuthRedirectURL(ctx context.Context, provider string) (string, error)
	HandleOAuthCallback(ctx context.Context, provider, code, state string) (string, error)
} 