package entity

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	AccountId   string    `gorm:"column:accountId;primary_key;default:''" json:"accountId"` // 账户ID
	Username    string    `gorm:"column:username;default:''" json:"username"`               // 用户名
	Password    string    `gorm:"column:password;default:''" json:"password"`               // 密码
	Secret      string    `gorm:"column:secret;default:''" json:"secret"`                   // 密钥
	Type        int       `gorm:"column:type;default:0" json:"type"`                        // 账户类型 0:管理员 1:普通用户
	CompanyCode string    `gorm:"column:companyCode;default:''" json:"companyCode"`         // 公司代码
	Status      int       `gorm:"column:status;default:0" json:"status"`                    // 状态 0:启用 1:禁用
	WhiteIp     string    `gorm:"column:whiteIp;default:''" json:"whiteIp"`                 // 白名单IP，多个IP用逗号分隔。用于限制IP访问
	CreateTime  time.Time `gorm:"column:createTime;default:''" json:"createTime"`           // 创建时间
	UpdateTime  time.Time `gorm:"column:updateTime;default:''" json:"updateTime"`           // 更新时间
}

func (c *Account) TableName() string {
	return "t_account"
}

// 默认填充创建时间
func (c *Account) BeforeCreate(tx *gorm.DB) error {
	c.CreateTime = time.Now()
	c.UpdateTime = time.Now()
	return nil
}

// 默认填充更新时间
func (c *Account) BeforeUpdate(tx *gorm.DB) error {
	c.UpdateTime = time.Now()
	return nil
}
