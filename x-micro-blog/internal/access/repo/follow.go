package repo

import (
	"x-micro-blog/internal/access/model"

	"gorm.io/gorm"
)

// FollowRepository 关注关系仓库接口
type FollowRepository interface {
	// Create 创建关注关系
	Create(follow *model.Follow) error

	// Delete 删除关注关系
	Delete(userID, targetID uint64) error

	// Exists 检查关注关系是否存在
	Exists(userID, targetID uint64) (bool, error)

	// ListFollowers 获取用户的粉丝列表
	ListFollowers(userID uint64, offset, limit int) ([]*model.User, int64, error)

	// ListFollowing 获取用户的关注列表
	ListFollowing(userID uint64, offset, limit int) ([]*model.User, int64, error)

	// CountFollowers 获取用户���粉丝数量
	CountFollowers(userID uint64) (int64, error)

	// CountFollowing 获取用户的关注数量
	CountFollowing(userID uint64) (int64, error)
}

// FollowRepositoryImpl 关注关系仓库实现
type FollowRepositoryImpl struct {
	db *gorm.DB
}

// NewFollowRepository 创建关注关系仓库
func NewFollowRepository(db *gorm.DB) FollowRepository {
	return &FollowRepositoryImpl{db: db}
}

// Create 创建关注关系
func (r *FollowRepositoryImpl) Create(follow *model.Follow) error {
	return r.db.Create(follow).Error
}

// Delete 删除关注关系
func (r *FollowRepositoryImpl) Delete(userID, targetID uint64) error {
	return r.db.Where("user_id = ? AND target_id = ?", userID, targetID).Delete(&model.Follow{}).Error
}

// Exists 检查关注关系是否存在
func (r *FollowRepositoryImpl) Exists(userID, targetID uint64) (bool, error) {
	var count int64
	err := r.db.Model(&model.Follow{}).Where("user_id = ? AND target_id = ?", userID, targetID).Count(&count).Error
	return count > 0, err
}

// ListFollowers 获取用户的粉丝列表
func (r *FollowRepositoryImpl) ListFollowers(userID uint64, offset, limit int) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64

	// 获取总数
	err := r.db.Model(&model.Follow{}).Where("target_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取粉丝列表
	err = r.db.Model(&model.Follow{}).
		Select("users.*").
		Joins("LEFT JOIN users ON users.id = follows.user_id").
		Where("follows.target_id = ?", userID).
		Offset(offset).
		Limit(limit).
		Find(&users).Error

	return users, total, err
}

// ListFollowing 获取用户的关注列表
func (r *FollowRepositoryImpl) ListFollowing(userID uint64, offset, limit int) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64

	// 获取总数
	err := r.db.Model(&model.Follow{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取关注列表
	err = r.db.Model(&model.Follow{}).
		Select("users.*").
		Joins("LEFT JOIN users ON users.id = follows.target_id").
		Where("follows.user_id = ?", userID).
		Offset(offset).
		Limit(limit).
		Find(&users).Error

	return users, total, err
}

// CountFollowers 获取用户的粉丝数量
func (r *FollowRepositoryImpl) CountFollowers(userID uint64) (int64, error) {
	var count int64
	err := r.db.Model(&model.Follow{}).Where("target_id = ?", userID).Count(&count).Error
	return count, err
}

// CountFollowing 获取用户的关注数量
func (r *FollowRepositoryImpl) CountFollowing(userID uint64) (int64, error) {
	var count int64
	err := r.db.Model(&model.Follow{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}
