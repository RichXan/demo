package model

import (
	"database/sql/driver"
	"encoding/json"
)

// SocialAccount 社交账号模型
type SocialAccount struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	UserID   uint64 `gorm:"not null" json:"user_id"`
	Provider string `gorm:"not null" json:"provider"` // 提供商：github, google, wechat, qq, weibo
	OpenID   string `gorm:"not null" json:"open_id"`  // 第三方用户ID
	UnionID  string `json:"union_id,omitempty"`       // 第三方用户统一ID（如微信）
	Nickname string `json:"nickname,omitempty"`       // 第三方昵称
	Avatar   string `json:"avatar,omitempty"`         // 第三方头像
	Extra    JSON   `gorm:"type:json" json:"extra"`   // 额外信息
	User     *User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	PublicTime
}

// JSON 自定义JSON类型
type JSON map[string]interface{}

// Value 实现 driver.Valuer 接口
func (j JSON) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// Scan 实现 sql.Scanner 接口
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	return json.Unmarshal(value.([]byte), j)
}

// ExtraData 获取额外数据
func (s *SocialAccount) ExtraData() JSON {
	if s.Extra == nil {
		return make(JSON)
	}
	return s.Extra
}

// SetExtraData 设置额外数据
func (s *SocialAccount) SetExtraData(data interface{}) error {
	if data == nil {
		s.Extra = nil
		return nil
	}

	// 如果已经是JSON类型，直接赋值
	if j, ok := data.(JSON); ok {
		s.Extra = j
		return nil
	}

	// 否则尝试通过JSON编解码转换
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	var j JSON
	if err := json.Unmarshal(b, &j); err != nil {
		return err
	}

	s.Extra = j
	return nil
}

// TableName 表名
func (SocialAccount) TableName() string {
	return "social_account"
}
