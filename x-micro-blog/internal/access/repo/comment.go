package repo

import (
	"x-micro-blog/internal/access/model"

	"gorm.io/gorm"
)

// CommentRepository 评论仓库接口
type CommentRepository interface {
	Create(comment *model.Comment) error
	Update(comment *model.Comment) error
	Delete(id uint64) error
	FindByID(id uint64) (*model.Comment, error)
	ListByPostID(postID uint64, offset, limit int) ([]*model.Comment, int64, error)
	ListByUserID(userID uint64, offset, limit int) ([]*model.Comment, int64, error)
	ListReplies(commentID uint64, offset, limit int) ([]*model.Comment, int64, error)
}

// CommentRepositoryImpl 评论仓库实现
type CommentRepositoryImpl struct {
	db *gorm.DB
}

// NewCommentRepository 创建评论仓库
func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &CommentRepositoryImpl{db: db}
}

// Create 创建评论
func (r *CommentRepositoryImpl) Create(comment *model.Comment) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// 创建评论
		if err := tx.Create(comment).Error; err != nil {
			return err
		}

		// 更新文章评论数
		if err := tx.Model(&model.Post{}).Where("id = ?", comment.PostID).
			UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}

// Update 更新评论
func (r *CommentRepositoryImpl) Update(comment *model.Comment) error {
	return r.db.Save(comment).Error
}

// Delete 删除评论
func (r *CommentRepositoryImpl) Delete(id uint64) error {
	var comment model.Comment
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// 查找评论
		if err := tx.First(&comment, id).Error; err != nil {
			return err
		}

		// 软删除评论
		if err := tx.Model(&comment).Update("status", 0).Error; err != nil {
			return err
		}

		// 更新文章评论数
		if err := tx.Model(&model.Post{}).Where("id = ?", comment.PostID).
			UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}

// FindByID 根据ID查找评论
func (r *CommentRepositoryImpl) FindByID(id uint64) (*model.Comment, error) {
	var comment model.Comment
	err := r.db.Preload("User").Preload("Post").Preload("Parent").
		First(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// ListByPostID 获取文章的评论列表
func (r *CommentRepositoryImpl) ListByPostID(postID uint64, offset, limit int) ([]*model.Comment, int64, error) {
	var comments []*model.Comment
	var total int64

	err := r.db.Model(&model.Comment{}).
		Where("post_id = ? AND parent_id IS NULL", postID).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("User").Preload("Replies.User").
		Where("post_id = ? AND parent_id IS NULL", postID).
		Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&comments).Error
	if err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}

// ListByUserID 获取用户的评论列表
func (r *CommentRepositoryImpl) ListByUserID(userID uint64, offset, limit int) ([]*model.Comment, int64, error) {
	var comments []*model.Comment
	var total int64

	err := r.db.Model(&model.Comment{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("User").Preload("Post").
		Where("user_id = ?", userID).
		Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&comments).Error
	if err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}

// ListReplies 获取评论的回复列表
func (r *CommentRepositoryImpl) ListReplies(commentID uint64, offset, limit int) ([]*model.Comment, int64, error) {
	var comments []*model.Comment
	var total int64

	err := r.db.Model(&model.Comment{}).Where("parent_id = ?", commentID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("User").
		Where("parent_id = ?", commentID).
		Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&comments).Error
	if err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}
