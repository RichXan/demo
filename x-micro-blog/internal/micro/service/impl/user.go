package impl

import (
	"context"
	"errors"
	"regexp"

	"x-micro-blog/internal/access/model"
	"x-micro-blog/internal/access/repo"
	"x-micro-blog/internal/micro/service"

	"github.com/RichXan/xcommon/xerror"
	"github.com/RichXan/xcommon/xlog"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// userService 用户服务实现
type userService struct {
	db         *gorm.DB
	logger     *xlog.Logger
	userRepo   repo.UserRepository
	followRepo repo.FollowRepository
}

// NewUserService 创建用户服务
func NewUserService(db *gorm.DB, logger *xlog.Logger) service.UserService {
	return &userService{
		db:         db,
		logger:     logger,
		userRepo:   repo.NewUserRepository(db),
		followRepo: repo.NewFollowRepository(db),
	}
}

// validateEmail 验证邮箱格式
func validateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

// validateUsername 验证用户名格式
func validateUsername(username string) bool {
	pattern := `^[a-zA-Z0-9_-]{4,16}$`
	match, _ := regexp.MatchString(pattern, username)
	return match
}

// Register 用户注册
func (s *userService) Register(ctx context.Context, username, password, email string) (*model.User, error) {
	// 验证用户名格式
	if !validateUsername(username) {
		return nil, xerror.UsernameInvalid
	}

	// 验证邮箱格式
	if !validateEmail(email) {
		return nil, xerror.EmailInvalid
	}

	// 检查用户名是否已存在
	if _, err := s.userRepo.FindByUsername(username); err == nil {
		return nil, xerror.UserExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerror.Wrap(err, xerror.CodeSystemError, "failed to check username")
	}

	// 检查邮箱是否已存在
	if _, err := s.userRepo.FindByEmail(email); err == nil {
		return nil, xerror.EmailExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerror.Wrap(err, xerror.CodeSystemError, "failed to check email")
	}

	// 创建用户
	user := &model.User{
		Username: username,
		Password: password,
		Email:    email,
		Nickname: username,
		Status:   1,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, xerror.Wrap(err, xerror.CodeSystemError, "failed to create user")
	}

	return user, nil
}

// Login 用户登录
func (s *userService) Login(ctx context.Context, username, password string) (string, error) {
	// 查找用户
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", xerror.UserNotFound
		}
		return "", xerror.Wrap(err, xerror.CodeSystemError, "failed to find user")
	}

	// 检查用户状态
	if user.Status != 1 {
		return "", xerror.UserDisabled
	}

	// 验证密码
	if !user.ComparePassword(password) {
		return "", xerror.PasswordError
	}

	// TODO: 生成访问令牌
	return "", nil
}

// GetUser 获取用户信息
func (s *userService) GetUser(ctx context.Context, userID int64) (*model.User, error) {
	user, err := s.userRepo.FindByID(uint64(userID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerror.UserNotFound
		}
		return nil, xerror.Wrap(err, xerror.CodeSystemError, "failed to get user")
	}
	return user, nil
}

// UpdateUser ��新用户信息
func (s *userService) UpdateUser(ctx context.Context, userID int64, nickname, avatar, bio string) (*model.User, error) {
	user, err := s.userRepo.FindByID(uint64(userID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerror.UserNotFound
		}
		return nil, xerror.Wrap(err, xerror.CodeSystemError, "failed to get user")
	}

	// 更新用户信息
	user.Nickname = nickname
	user.Avatar = avatar
	// TODO: 添加 bio 字段

	if err := s.userRepo.Update(user); err != nil {
		return nil, xerror.Wrap(err, xerror.CodeSystemError, "failed to update user")
	}

	return user, nil
}

// ChangePassword 修改密码
func (s *userService) ChangePassword(ctx context.Context, userID int64, oldPassword, newPassword string) error {
	user, err := s.userRepo.FindByID(uint64(userID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.UserNotFound
		}
		return xerror.Wrap(err, xerror.CodeSystemError, "failed to get user")
	}

	// 验证旧密码
	if !user.ComparePassword(oldPassword) {
		return xerror.PasswordError
	}

	// 生成新密码哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return xerror.Wrap(err, xerror.CodeSystemError, "failed to hash password")
	}

	// 更新密码
	user.Password = string(hashedPassword)
	if err := s.userRepo.Update(user); err != nil {
		return xerror.Wrap(err, xerror.CodeSystemError, "failed to update password")
	}

	return nil
}

// ListUsers 获取用户列表
func (s *userService) ListUsers(ctx context.Context, page, pageSize int) ([]*model.User, int64, error) {
	offset := (page - 1) * pageSize
	return s.userRepo.List(offset, pageSize)
}

// ListSocialAccounts 获取社交账号列表
func (s *userService) ListSocialAccounts(ctx context.Context, userID int64) ([]*model.SocialAccount, error) {
	accounts, err := s.userRepo.ListSocialAccounts(uint64(userID))
	if err != nil {
		return nil, xerror.Wrap(err, xerror.CodeSystemError, "failed to list social accounts")
	}

	result := make([]*model.SocialAccount, len(accounts))
	for i := range accounts {
		result[i] = &accounts[i]
	}
	return result, nil
}

// BindSocialAccount 绑定社交账号
func (s *userService) BindSocialAccount(ctx context.Context, userID int64, provider, code, state string) error {
	// TODO: 实现社交账号绑定逻辑
	return nil
}

// UnbindSocialAccount 解绑社交账号
func (s *userService) UnbindSocialAccount(ctx context.Context, userID int64, provider string) error {
	accounts, err := s.userRepo.ListSocialAccounts(uint64(userID))
	if err != nil {
		return xerror.Wrap(err, xerror.CodeSystemError, "failed to list social accounts")
	}

	// 检查是否是最后一个社交账号
	if len(accounts) == 1 && accounts[0].Provider == provider {
		return xerror.NewError(xerror.CodeRequestRejected, "cannot unbind the last social account")
	}

	// 查找要解绑的社交账号
	var accountID uint64
	for _, account := range accounts {
		if account.Provider == provider {
			accountID = account.ID
			break
		}
	}

	if accountID == 0 {
		return xerror.NewError(xerror.CodeNotFound, "social account not found")
	}

	// 删除社交账号
	if err := s.userRepo.DeleteSocialAccount(accountID); err != nil {
		return xerror.Wrap(err, xerror.CodeSystemError, "failed to delete social account")
	}

	return nil
}

// FollowUser 关注用户
func (s *userService) FollowUser(ctx context.Context, userID, targetID int64) error {
	// 检查目标用户是否存在
	_, err := s.userRepo.FindByID(uint64(targetID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.NewError(xerror.CodeNotFound, "target user not found")
		}
		return xerror.Wrap(err, xerror.CodeSystemError, "failed to find target user")
	}

	// 不能关注自己
	if userID == targetID {
		return xerror.NewError(xerror.CodeParamError, "cannot follow yourself")
	}

	// 检查是否已经关注
	exists, err := s.followRepo.Exists(uint64(userID), uint64(targetID))
	if err != nil {
		return xerror.Wrap(err, xerror.CodeSystemError, "failed to check follow status")
	}
	if exists {
		return xerror.NewError(xerror.CodeParamError, "already following this user")
	}

	// 创建关注关系
	follow := &model.Follow{
		UserID:   uint64(userID),
		TargetID: uint64(targetID),
	}
	if err := s.followRepo.Create(follow); err != nil {
		return xerror.Wrap(err, xerror.CodeSystemError, "failed to create follow relationship")
	}

	return nil
}

// UnfollowUser 取消关注
func (s *userService) UnfollowUser(ctx context.Context, userID, targetID int64) error {
	// 检查是否已经关注
	exists, err := s.followRepo.Exists(uint64(userID), uint64(targetID))
	if err != nil {
		return xerror.Wrap(err, xerror.CodeSystemError, "failed to check follow status")
	}
	if !exists {
		return xerror.NewError(xerror.CodeParamError, "not following this user")
	}

	// 删除关注关系
	if err := s.followRepo.Delete(uint64(userID), uint64(targetID)); err != nil {
		return xerror.Wrap(err, xerror.CodeSystemError, "failed to delete follow relationship")
	}

	return nil
}

// ListFollowers 获取粉丝列表
func (s *userService) ListFollowers(ctx context.Context, userID int64, page, pageSize int) ([]*model.User, int64, error) {
	offset := (page - 1) * pageSize
	return s.followRepo.ListFollowers(uint64(userID), offset, pageSize)
}

// ListFollowing 获取关注列表
func (s *userService) ListFollowing(ctx context.Context, userID int64, page, pageSize int) ([]*model.User, int64, error) {
	offset := (page - 1) * pageSize
	return s.followRepo.ListFollowing(uint64(userID), offset, pageSize)
}
