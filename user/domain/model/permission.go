package model

type Permission struct {
	ID                 int64  `gorm:"primary_key;not_null;auto_increment"`
	PermissionName     string `json:"permission_name"`
	PermissionDescribe string `json:"permission_describe"`
	PermissionAction   string `json:"permission_action"`
	PermissionStatus   int32  `json:"permission_status"`
}
