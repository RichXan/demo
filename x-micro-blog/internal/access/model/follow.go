package model

import (
	"time"

	"gorm.io/gorm"
)

// Follow 关注关系模型
type Follow struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	UserID    uint64    `gorm:"not null" json:"user_id"`   // 关注者ID
	TargetID  uint64    `gorm:"not null" json:"target_id"` // 被关注者ID
	User      *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`     // 关注者
	Target    *User     `gorm:"foreignKey:TargetID" json:"target,omitempty"` // 被关注者
	PublicTime
}

func (t *Follow) TableName() string {
	return "follows"
}

func (t *Follow) BeforeCreate(tx *gorm.DB) (err error) {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return
}

func (t *Follow) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now()
	return
}
