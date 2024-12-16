// Package po generated by 'freedom new-po'
package po

import (
	"gorm.io/gorm"
	"time"
)

// Admin .
type Admin struct {
	changes map[string]interface{}
	ID      int       `gorm:"primaryKey;column:id"`
	Age     int       `gorm:"column:age"`
	RoleID  int       `gorm:"column:role_id"`
	Name    string    `gorm:"column:name"`
	Address string    `gorm:"column:address"`
	Created time.Time `gorm:"column:created"`
	Updated time.Time `gorm:"column:updated"`
}

// TableName .
func (obj *Admin) TableName() string {
	return "admin"
}

// Location .
func (obj *Admin) Location() map[string]interface{} {
	return map[string]interface{}{"id": obj.ID}
}

// GetChanges .
func (obj *Admin) GetChanges() map[string]interface{} {
	if obj.changes == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range obj.changes {
		result[k] = v
	}
	obj.changes = nil
	return result
}

// Update .
func (obj *Admin) Update(name string, value interface{}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.changes[name] = value
}

// SetAge .
func (obj *Admin) SetAge(age int) {
	obj.Age = age
	obj.Update("age", age)
}

// SetRoleID .
func (obj *Admin) SetRoleID(roleID int) {
	obj.RoleID = roleID
	obj.Update("role_id", roleID)
}

// SetName .
func (obj *Admin) SetName(name string) {
	obj.Name = name
	obj.Update("name", name)
}

// SetAddress .
func (obj *Admin) SetAddress(address string) {
	obj.Address = address
	obj.Update("address", address)
}

// SetCreated .
func (obj *Admin) SetCreated(created time.Time) {
	obj.Created = created
	obj.Update("created", created)
}

// SetUpdated .
func (obj *Admin) SetUpdated(updated time.Time) {
	obj.Updated = updated
	obj.Update("updated", updated)
}

// AddAge .
func (obj *Admin) AddAge(age int) {
	obj.Age += age
	obj.Update("age", gorm.Expr("age + ?", age))
}

// AddRoleID .
func (obj *Admin) AddRoleID(roleID int) {
	obj.RoleID += roleID
	obj.Update("role_id", gorm.Expr("role_id + ?", roleID))
}