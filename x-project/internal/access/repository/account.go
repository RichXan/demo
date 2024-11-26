package repository

import (
	"fmt"
	"slices"
	"strings"
	"xproject/internal/access/entity"

	"gorm.io/gorm"
)

type AccountRepository interface {
	/*
		Create 在数据库事务中创建一个新的 Account 记录。

		参数:
		- dbTx *gorm.DB: 表示一个数据库事务实例，用于执行创建操作。
		- account *entity.Account: 指向一个 Account 结构体的指针，包含了将要被创建的记录的信息。

		返回值:
		- error: 如果创建过程中遇到任何错误，则返回错误信息；否则返回 nil。
	*/
	Create(account *entity.Account) error
	/*
		Delete 删除一个 Account 记录。
	*/
	Delete(accountId string) error
	/*
		Update 更新一个 Account 记录。
	*/
	Update(account *entity.Account) error
	/*
		Get 根据账户ID获取一个 Account 记录。
	*/
	Get(accountId string) (*entity.Account, error)
	/*
		Login 根据账户ID、密钥和IP地址登录。
	*/
	Login(accountId, secret, ip string) (*entity.Account, error)
}

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) AccountRepository {
	return &accountRepo{
		db: db,
	}
}

func (r *accountRepo) Create(account *entity.Account) error {
	return r.db.Create(account).Error
}

func (r *accountRepo) Delete(accountId string) error {
	return r.db.Delete(&entity.Account{}, accountId).Error
}

func (r *accountRepo) Update(account *entity.Account) error {
	return r.db.Save(account).Error
}

func (r *accountRepo) Get(accountId string) (*entity.Account, error) {
	var account entity.Account
	if err := r.db.Where("accountId = ?", accountId).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepo) Login(accountId, secret, ip string) (*entity.Account, error) {
	var account entity.Account
	if err := r.db.Where("accountId = ?", accountId).First(&account).Error; err != nil {
		return nil, err
	}
	if account.Status != 0 {
		return nil, fmt.Errorf("account is inactive")
	}
	if account.Secret != secret {
		return nil, fmt.Errorf("accountId or secret is incorrect")
	}

	// 白名单校验, 白名单为空则不校验, 否则校验ip是否在白名单中
	if account.WhiteIp != "" && !slices.Contains(strings.Split(account.WhiteIp, ","), ip) {
		return nil, fmt.Errorf("ip is not in the white list")
	}
	return &account, nil
}
