package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yunixiangfeng/gopaas/user/domain/model"
)

//创建需要实现的接口
type IUserRepository interface {
	//初始化表
	InitTable() error
	//根据ID查处找数据
	FindUserByID(int64) (*model.User, error)
	//创建一条 user 数据
	CreateUser(*model.User) (int64, error)
	//根据ID删除一条 user 数据
	DeleteUserByID(int64) error
	//修改更新数据
	UpdateUser(*model.User) error
	//查找user所有数据
	FindAll() ([]model.User, error)

	// 分配角色
	AddRole(*model.User, []*model.Role) error
	// 更新用户角色
	UpdateRole(*model.User, []*model.Role) error
	// 删除用户的角色
	DeleteRole(*model.User, []*model.Role) error
	// 判断用户是否有对应的权限
	IsRight(string, int64) bool
}

//创建userRepository
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

// 为已经存在的用户添加角色
func (u *UserRepository) AddRole(user *model.User, roles []*model.Role) error {
	return u.mysqlDb.Model(&user).Association("Role").Append(roles).Error
}

// 更新用户角色
func (u *UserRepository) UpdateRole(user *model.User, roles []*model.Role) error {
	return u.mysqlDb.Model(&user).Association("Role").Replace(roles).Error
}

// 删除用户角色
func (u *UserRepository) DeleteRole(user *model.User, roles []*model.Role) error {
	return u.mysqlDb.Model(&user).Association("Role").Delete(roles).Error
}

// 检测当前用户是否具备权限
func (u *UserRepository) IsRight(action string, userID int64) bool {
	permission := &model.Permission{}
	sql := "select p.id From user u,user_role ur,role r,role_permission rp, permission p WHERE p.permission_action=? AND p.id = rp.permission_id AND rp.role_id=r,id AND ur.role_id AND ur.user_id=u.id AND u.id=?"
	u.mysqlDb.Raw(sql, action, userID).Scan(permission)
	// 可以写其它判断逻辑
	if permission.ID > 0 {
		return true
	}
	return false
}

//初始化表
func (u *UserRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.User{}, &model.Role{}, &model.Permission{}).Error
}

//根据ID查找User信息
func (u *UserRepository) FindUserByID(userID int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.First(user, userID).Error
}

//创建User信息
func (u *UserRepository) CreateUser(user *model.User) (int64, error) {
	return user.ID, u.mysqlDb.Create(user).Error
}

//根据ID删除User信息
func (u *UserRepository) DeleteUserByID(userID int64) error {
	return u.mysqlDb.Where("id = ?", userID).Delete(&model.User{}).Error
}

//更新User信息
func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Update(user).Error
}

//获取结果集
func (u *UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.mysqlDb.Find(&userAll).Error
}
