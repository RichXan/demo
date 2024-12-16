package repo

import (
	"x-micro-blog/internal/access/model"

	"gorm.io/gorm"
)

// LikeRepository 点赞仓库接口
type LikeRepository interface {
	Create(like *model.Like) error
	Delete(userID, postID uint64) error
	Exists(userID, postID uint64) (bool, error)
	ListByPostID(postID uint64, offset, limit int) ([]*model.Like, int64, error)
	ListByUserID(userID uint64, offset, limit int) ([]*model.Like, int64, error)
}

// LikeRepositoryImpl 点赞仓库实现
type LikeRepositoryImpl struct {
	db *gorm.DB
}

// NewLikeRepository 创建点赞仓库
func NewLikeRepository(db *gorm.DB) LikeRepository {
	return &LikeRepositoryImpl{db: db}
}

// Create 创建点赞
func (r *LikeRepositoryImpl) Create(like *model.Like) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// 创建点赞记录
		if err := tx.Create(like).Error; err != nil {
			return err
		}

		// 更新文章点赞数
		if err := tx.Model(&model.Post{}).Where("id = ?", like.PostID).
			UpdateColumn("like_count", gorm.Expr("like_count + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}

// Delete 删除点赞
func (r *LikeRepositoryImpl) Delete(userID, postID uint64) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// 删除点赞记录
		result := tx.Where("user_id = ? AND post_id = ?", userID, postID).Delete(&model.Like{})
		if result.Error != nil {
			return result.Error
		}

		// 如果确实删除了记录，则更新文章点赞数
		if result.RowsAffected > 0 {
			if err := tx.Model(&model.Post{}).Where("id = ?", postID).
				UpdateColumn("like_count", gorm.Expr("like_count - ?", 1)).Error; err != nil {
				return err
			}
		}

		return nil
	})
	return err
}

// Exists 检查点赞是否存在
func (r *LikeRepositoryImpl) Exists(userID, postID uint64) (bool, error) {
	var count int64
	err := r.db.Model(&model.Like{}).
		Where("user_id = ? AND post_id = ?", userID, postID).
		Count(&count).Error
	return count > 0, err
}

// ListByPostID 获取文章的点赞列表
func (r *LikeRepositoryImpl) ListByPostID(postID uint64, offset, limit int) ([]*model.Like, int64, error) {
	var likes []*model.Like
	var total int64

	err := r.db.Model(&model.Like{}).Where("post_id = ?", postID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("User").
		Where("post_id = ?", postID).
		Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&likes).Error
	if err != nil {
		return nil, 0, err
	}

	return likes, total, nil
}

// ListByUserID 获取用户的点赞列表
func (r *LikeRepositoryImpl) ListByUserID(userID uint64, offset, limit int) ([]*model.Like, int64, error) {
	var likes []*model.Like
	var total int64

	err := r.db.Model(&model.Like{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("Post").
		Where("user_id = ?", userID).
		Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&likes).Error
	if err != nil {
		return nil, 0, err
	}

	return likes, total, nil
}
