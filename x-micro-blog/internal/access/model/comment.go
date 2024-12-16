package model

import (
	"time"

	"gorm.io/gorm"
)

// Comment 评论模型
type Comment struct {
	ID       uint64    `gorm:"primaryKey" json:"id"`
	PostID   uint64    `gorm:"not null" json:"post_id"`
	UserID   uint64    `gorm:"not null" json:"user_id"`
	Content  string    `gorm:"type:text;not null" json:"content"`
	ParentID *uint64   `json:"parent_id"`
	Status   int8      `gorm:"default:1;not null" json:"status"` // 0-已删除，1-正常
	User     *User     `gorm:"foreignKey:UserID" json:"user"`
	Post     *Post     `gorm:"foreignKey:PostID" json:"post"`
	Parent   *Comment  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Replies  []Comment `gorm:"foreignKey:ParentID" json:"replies,omitempty"`
	PublicTime
}

func (t *Comment) TableName() string {
	return "comments"
}

func (t *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return
}

func (t *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now()
	return
}
