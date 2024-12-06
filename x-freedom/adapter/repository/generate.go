package repository

import (
	"errors"
	"fmt"
	"github.com/8treenet/freedom"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
	"time"
	"x-freedom/domain/po"
)

// GORMRepository .
type GORMRepository interface {
	db() *gorm.DB
	Worker() freedom.Worker
}

type saveObject interface {
	TableName() string
	Location() map[string]interface{}
	GetChanges() map[string]interface{}
}

// Builder .
type Builder interface {
	Execute(db *gorm.DB, object interface{}) error
}

// Pager .
type Pager struct {
	pageSize  int
	page      int
	totalPage int
	fields    []string
	orders    []string
}

// NewDescPager .
func NewDescPager(column string, columns ...string) *Pager {
	return newDefaultPager("desc", column, columns...)
}

// NewAscPager .
func NewAscPager(column string, columns ...string) *Pager {
	return newDefaultPager("asc", column, columns...)
}

// NewDescOrder .
func newDefaultPager(sort, field string, args ...string) *Pager {
	fields := []string{field}
	fields = append(fields, args...)
	orders := []string{}
	for index := 0; index < len(fields); index++ {
		orders = append(orders, sort)
	}
	return &Pager{
		fields: fields,
		orders: orders,
	}
}

// Order .
func (p *Pager) Order() interface{} {
	if len(p.fields) == 0 {
		return nil
	}
	args := []string{}
	for index := 0; index < len(p.fields); index++ {
		args = append(args, fmt.Sprintf("`%s` %s", p.fields[index], p.orders[index]))
	}

	return strings.Join(args, ",")
}

// TotalPage .
func (p *Pager) TotalPage() int {
	return p.totalPage
}

// SetPage .
func (p *Pager) SetPage(page, pageSize int) *Pager {
	p.page = page
	p.pageSize = pageSize
	return p
}

// Execute .
func (p *Pager) Execute(db *gorm.DB, object interface{}) (e error) {
	if p.page != 0 && p.pageSize != 0 {
		var count64 int64
		e = db.Model(object).Count(&count64).Error
		count := int(count64)
		if e != nil {
			return
		}
		if count != 0 {
			//Calculate the length of the pagination
			if count%p.pageSize == 0 {
				p.totalPage = count / p.pageSize
			} else {
				p.totalPage = count/p.pageSize + 1
			}
		}
		db = db.Offset((p.page - 1) * p.pageSize).Limit(p.pageSize)
	}

	orderValue := p.Order()
	if orderValue != nil {
		db = db.Order(orderValue)
	}

	resultDB := db.Find(object)
	if resultDB.Error != nil {
		return resultDB.Error
	}
	return
}

// Limiter .
type Limiter struct {
	size   int
	column string
	desc   bool
}

// NewDescLimiter .
func NewDescLimiter(column string, size int) *Limiter {
	return &Limiter{column: column, size: size, desc: true}
}

// NewAscLimiter .
func NewAscLimiter(column string, size int) *Limiter {
	return &Limiter{column: column, size: size, desc: false}
}

// Execute .
func (limiter *Limiter) Execute(db *gorm.DB, object interface{}) (e error) {
	db = db.Order(clause.OrderByColumn{Column: clause.Column{Name: limiter.column}, Desc: limiter.desc}).Limit(limiter.size)
	resultDB := db.Find(object)
	if resultDB.Error != nil {
		return resultDB.Error
	}
	return
}

func ormErrorLog(repo GORMRepository, model, method string, e error, expression ...interface{}) {
	if e == nil || e == gorm.ErrRecordNotFound {
		return
	}
	repo.Worker().Logger().Errorf("error: %v, model: %s, method: %s", e, model, method)
}

// findAccount .
func findAccount(repo GORMRepository, result *po.Account, builders ...Builder) (e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Account", "findAccount", e, now)
		ormErrorLog(repo, "Account", "findAccount", e, result)
	}()
	db := repo.db()
	if len(builders) == 0 {
		e = db.Where(result).Last(result).Error
		return
	}
	e = builders[0].Execute(db.Limit(1), result)
	return
}

// findAccountByUID .
func findAccountByUID(repo GORMRepository, uID int) (result po.Account, e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Account", "findAccountUID", e, now)
		ormErrorLog(repo, "Account", "findAccountUID", e, uID)
	}()

	e = repo.db().Last(&result, uID).Error
	return
}

// findAccountListByUID .
func findAccountListByUID(repo GORMRepository, uID ...int) (result []*po.Account, e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Account", "findAccountListByUID", e, now)
		ormErrorLog(repo, "Account", "findAccountListByUID", e, uID)
	}()

	e = repo.db().Find(&result, uID).Error
	return
}

// findAccountByWhere .
func findAccountByWhere(repo GORMRepository, query string, args []interface{}, builders ...Builder) (result po.Account, e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Account", "findAccountByWhere", e, now)
		ormErrorLog(repo, "Account", "findAccountByWhere", e, query, args)
	}()
	db := repo.db()
	if query != "" {
		db = db.Where(query, args...)
	}
	if len(builders) == 0 {
		e = db.Last(&result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), &result)
	return
}

// findAccountByMap .
func findAccountByMap(repo GORMRepository, query map[string]interface{}, builders ...Builder) (result po.Account, e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Account", "findAccountByMap", e, now)
		ormErrorLog(repo, "Account", "findAccountByMap", e, query)
	}()

	db := repo.db().Where(query)
	if len(builders) == 0 {
		e = db.Last(&result).Error
		return
	}

	e = builders[0].Execute(db.Limit(1), &result)
	return
}

// findAccountList .
func findAccountList(repo GORMRepository, query po.Account, builders ...Builder) (results []*po.Account, e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Account", "findAccountList", e, now)
		ormErrorLog(repo, "Account", "findAccounts", e, query)
	}()
	db := repo.db().Where(query)

	if len(builders) == 0 {
		e = db.Find(&results).Error
		return
	}
	e = builders[0].Execute(db, &results)
	return
}

// findAccountListByWhere .
func findAccountListByWhere(repo GORMRepository, query string, args []interface{}, builders ...Builder) (results []*po.Account, e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Account", "findAccountListByWhere", e, now)
		ormErrorLog(repo, "Account", "findAccountsByWhere", e, query, args)
	}()
	db := repo.db()
	if query != "" {
		db = db.Where(query, args...)
	}

	if len(builders) == 0 {
		e = db.Find(&results).Error
		return
	}
	e = builders[0].Execute(db, &results)
	return
}

// findAccountListByMap .
func findAccountListByMap(repo GORMRepository, query map[string]interface{}, builders ...Builder) (results []*po.Account, e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Account", "findAccountListByMap", e, now)
		ormErrorLog(repo, "Account", "findAccountsByMap", e, query)
	}()

	db := repo.db().Where(query)

	if len(builders) == 0 {
		e = db.Find(&results).Error
		return
	}
	e = builders[0].Execute(db, &results)
	return
}

// AccountListToMap
func AccountListToMap(list []*po.Account, inErr error) (result map[int]*po.Account, e error) {
	result = make(map[int]*po.Account)
	if inErr != nil {
		e = inErr
		return
	}
	for _, v := range list {
		result[v.UID] = v
	}
	return
}

// AccountToPoint
func AccountToPoint(object po.Account, inErr error) (result *po.Account, e error) {
	if inErr != nil {
		e = inErr
		return
	}
	result = &object
	return
}

// createAccount .
func createAccount(repo GORMRepository, object *po.Account) (rowsAffected int64, e error) {
	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Account", "createAccount", e, now)
		ormErrorLog(repo, "Account", "createAccount", e, *object)
	}()

	db := repo.db().Create(object)
	rowsAffected = db.RowsAffected
	e = db.Error
	return
}

// saveAccount .
func saveAccount(repo GORMRepository, object saveObject) (rowsAffected int64, e error) {
	if len(object.Location()) == 0 {
		return 0, errors.New("location cannot be empty")
	}
	updateValues := object.GetChanges()
	if len(updateValues) == 0 {
		return 0, nil
	}

	now := time.Now()
	defer func() {
		freedom.Prometheus().OrmWithLabelValues("Account", "saveAccount", e, now)
		ormErrorLog(repo, "Account", "saveAccount", e, object)
	}()

	db := repo.db().Table(object.TableName()).Where(object.Location()).Updates(updateValues)
	e = db.Error
	rowsAffected = db.RowsAffected
	return
}
