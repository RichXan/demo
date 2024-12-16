package model

import (
	"time"

	"gorm.io/gorm"
)

// Post 文章模型
type Post struct {
	ID           uint64 `gorm:"primaryKey" json:"id"`
	Title        string `gorm:"size:200;not null" json:"title"`
	Content      string `gorm:"type:text;not null" json:"content"`
	UserID       uint64 `gorm:"not null" json:"user_id"`
	Status       int8   `gorm:"default:1;not null" json:"status"` // 0-草稿，1-已发布，2-已删除
	ViewCount    uint32 `gorm:"default:0;not null" json:"view_count"`
	LikeCount    uint32 `gorm:"default:0;not null" json:"like_count"`
	CommentCount uint32 `gorm:"default:0;not null" json:"comment_count"`
	Tags         []Tag  `gorm:"many2many:post_tags;" json:"tags"`
	User         *User  `gorm:"foreignKey:UserID" json:"user"`
	PublicTime
}

func (t *Post) TableName() string {
	return "posts"
}

func (t *Post) BeforeCreate(tx *gorm.DB) (err error) {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return
}

func (t *Post) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now()
	return
}

// Tag 标签模型
type Tag struct {
	ID    uint64 `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"uniqueIndex;size:50;not null" json:"name"`
	Posts []Post `gorm:"many2many:post_tags;" json:"-"`
	PublicTime
}

func (t *Tag) TableName() string {
	return "tags"
}

func (t *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	t.CreatedAt = time.Now()
	return
}

func (t *Tag) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now()
	return
}
