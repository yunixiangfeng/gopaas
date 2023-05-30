package model

type Role struct {
	ID         int64         `gorm:"primary_key;not_null;auto_increment"`
	RoleName   string        `json:"role_name`
	RoleStatus int32         `json:"role_status"`
	Permission []*Permission `gorm:"many2many:role_permission"`
}
