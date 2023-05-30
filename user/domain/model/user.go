package model

type User struct {
	ID         int64   `gorm:"primary_key;not_null;auto_increment"`
	UserName   string  `gorm:"not_null;unique" json:"user_name"`
	UserEmail  string  `gorm:"not_null;unique" json:"user_email"`
	IsAdmin    bool    `json:"is_admin"`
	UserPwd    string  `json:"user_pwd`
	UserStatus int32   `json:"user_status"`
	Role       []*Role `gorm:"many2many:user_role"`
}
