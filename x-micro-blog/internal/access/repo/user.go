package repo

import (
	"x-micro-blog/internal/access/model"

	"gorm.io/gorm"
)

// UserRepository 用户仓库接口
type UserRepository interface {
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id uint64) error
	FindByID(id uint64) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	List(offset, limit int) ([]*model.User, int64, error)

	// 社交账号相关方法
	CreateSocialAccount(account *model.SocialAccount) error
	DeleteSocialAccount(id uint64) error
	FindSocialAccount(provider, openID string) (*model.SocialAccount, error)
	ListSocialAccounts(userID uint64) ([]model.SocialAccount, error)
}

// UserRepositoryImpl 用户仓库实现
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository 创建用户仓库
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

// Create 创建用户
func (r *UserRepositoryImpl) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// Update 更新用户
func (r *UserRepositoryImpl) Update(user *model.User) error {
	return r.db.Save(user).Error
}

// Delete 删除用户
func (r *UserRepositoryImpl) Delete(id uint64) error {
	return r.db.Delete(&model.User{}, id).Error
}

// FindByID 根据ID查找用户
func (r *UserRepositoryImpl) FindByID(id uint64) (*model.User, error) {
	var user model.User
	err := r.db.Preload("SocialAccounts").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByUsername 根据用户名查找用户
func (r *UserRepositoryImpl) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Preload("SocialAccounts").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail 根据邮箱查找用户
func (r *UserRepositoryImpl) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Preload("SocialAccounts").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// List 获取用户列表
func (r *UserRepositoryImpl) List(offset, limit int) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64

	err := r.db.Model(&model.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("SocialAccounts").Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// CreateSocialAccount 创建社交账号
func (r *UserRepositoryImpl) CreateSocialAccount(account *model.SocialAccount) error {
	return r.db.Create(account).Error
}

// DeleteSocialAccount 删除社交账号
func (r *UserRepositoryImpl) DeleteSocialAccount(id uint64) error {
	return r.db.Delete(&model.SocialAccount{}, id).Error
}

// FindSocialAccount 查找社交账号
func (r *UserRepositoryImpl) FindSocialAccount(provider, openID string) (*model.SocialAccount, error) {
	var account model.SocialAccount
	err := r.db.Where("provider = ? AND open_id = ?", provider, openID).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// ListSocialAccounts 获取用户的所有社交账号
func (r *UserRepositoryImpl) ListSocialAccounts(userID uint64) ([]model.SocialAccount, error) {
	var accounts []model.SocialAccount
	err := r.db.Where("user_id = ?", userID).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
