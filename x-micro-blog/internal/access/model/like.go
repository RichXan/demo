package model

import (
	"time"

	"gorm.io/gorm"
)

// Like 点赞模型
type Like struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	UserID    uint64    `gorm:"not null" json:"user_id"`
	PostID    uint64    `gorm:"not null" json:"post_id"`
	User      *User     `gorm:"foreignKey:UserID" json:"user"`
	Post      *Post     `gorm:"foreignKey:PostID" json:"post"`
	PublicTime
}

func (t *Like) TableName() string {
	return "likes"
}

func (t *Like) BeforeCreate(tx *gorm.DB) (err error) {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return
}

func (t *Like) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now()
	return
}
