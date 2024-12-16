package repo

import (
	"x-micro-blog/internal/access/model"

	"gorm.io/gorm"
)

// PostRepository 文章仓库接口
type PostRepository interface {
	Create(post *model.Post) error
	Update(post *model.Post) error
	Delete(id uint64) error
	FindByID(id uint64) (*model.Post, error)
	List(offset, limit int) ([]*model.Post, int64, error)
	ListByUserID(userID uint64, offset, limit int) ([]*model.Post, int64, error)
	ListByTag(tagName string, offset, limit int) ([]*model.Post, int64, error)
	IncrementViewCount(id uint64) error
	IncrementLikeCount(id uint64) error
	DecrementLikeCount(id uint64) error
	IncrementCommentCount(id uint64) error
	DecrementCommentCount(id uint64) error
}

// PostRepositoryImpl 文章仓库实现
type PostRepositoryImpl struct {
	db *gorm.DB
}

// NewPostRepository 创建文章仓库
func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{db: db}
}

// Create 创建文章
func (r *PostRepositoryImpl) Create(post *model.Post) error {
	return r.db.Create(post).Error
}

// Update 更新文章
func (r *PostRepositoryImpl) Update(post *model.Post) error {
	return r.db.Save(post).Error
}

// Delete 删除文章
func (r *PostRepositoryImpl) Delete(id uint64) error {
	return r.db.Delete(&model.Post{}, id).Error
}

// FindByID 根据ID查找文章
func (r *PostRepositoryImpl) FindByID(id uint64) (*model.Post, error) {
	var post model.Post
	err := r.db.Preload("User").Preload("Tags").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// List 获取文章列表
func (r *PostRepositoryImpl) List(offset, limit int) ([]*model.Post, int64, error) {
	var posts []*model.Post
	var total int64

	err := r.db.Model(&model.Post{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("User").Preload("Tags").
		Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

// ListByUserID 获取用户的文章列表
func (r *PostRepositoryImpl) ListByUserID(userID uint64, offset, limit int) ([]*model.Post, int64, error) {
	var posts []*model.Post
	var total int64

	err := r.db.Model(&model.Post{}).Where("user_id = ?", userID).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("User").Preload("Tags").
		Where("user_id = ?", userID).
		Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

// ListByTag 获取标签下的文章列表
func (r *PostRepositoryImpl) ListByTag(tagName string, offset, limit int) ([]*model.Post, int64, error) {
	var posts []*model.Post
	var total int64

	subQuery := r.db.Table("tags").
		Select("post_tags.post_id").
		Joins("JOIN post_tags ON tags.id = post_tags.tag_id").
		Where("tags.name = ?", tagName)

	err := r.db.Model(&model.Post{}).
		Where("id IN (?)", subQuery).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Preload("User").Preload("Tags").
		Where("id IN (?)", subQuery).
		Offset(offset).Limit(limit).
		Order("created_at DESC").
		Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

// IncrementViewCount 增加浏览次数
func (r *PostRepositoryImpl) IncrementViewCount(id uint64) error {
	return r.db.Model(&model.Post{}).Where("id = ?", id).
		UpdateColumn("view_count", gorm.Expr("view_count + ?", 1)).Error
}

// IncrementLikeCount 增加点赞次数
func (r *PostRepositoryImpl) IncrementLikeCount(id uint64) error {
	return r.db.Model(&model.Post{}).Where("id = ?", id).
		UpdateColumn("like_count", gorm.Expr("like_count + ?", 1)).Error
}

// DecrementLikeCount 减少点赞次数
func (r *PostRepositoryImpl) DecrementLikeCount(id uint64) error {
	return r.db.Model(&model.Post{}).Where("id = ?", id).
		UpdateColumn("like_count", gorm.Expr("like_count - ?", 1)).Error
}

// IncrementCommentCount 增加评论次数
func (r *PostRepositoryImpl) IncrementCommentCount(id uint64) error {
	return r.db.Model(&model.Post{}).Where("id = ?", id).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error
}

// DecrementCommentCount 减少评论次数
func (r *PostRepositoryImpl) DecrementCommentCount(id uint64) error {
	return r.db.Model(&model.Post{}).Where("id = ?", id).
		UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error
}
