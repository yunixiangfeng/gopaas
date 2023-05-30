package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yunixiangfeng/gopaas/user/domain/model"
)

//创建需要实现的接口
type IPermissionRepository interface {
	//根据ID查找permission数据
	FindPermissionByID(int64) (*model.Permission, error)
	//创建一条 permission 数据
	CreatePermission(*model.Permission) (int64, error)
	//根据ID删除一条 permission 数据
	DeletePermissionByID(int64) error
	//修改更新数据
	UpdatePermission(*model.Permission) error
	//查找permission所有数据
	FindAll() ([]model.Permission, error)

	// 根据ID查找所有权限
	FindAllPermissionById([]int64) ([]*model.Permission, error)
}

//创建PermissionRepository
func NewPermissionRepository(db *gorm.DB) IPermissionRepository {
	return &PermissionRepository{mysqlDb: db}
}

type PermissionRepository struct {
	mysqlDb *gorm.DB
}

//根据ID查找Permission信息
func (u *PermissionRepository) FindPermissionByID(permissionID int64) (permission *model.Permission, err error) {
	permission = &model.Permission{}
	return permission, u.mysqlDb.First(permission, permissionID).Error
}

//创建Permission信息
func (u *PermissionRepository) CreatePermission(permission *model.Permission) (int64, error) {
	return permission.ID, u.mysqlDb.Create(permission).Error
}

//根据ID删除Permission信息
func (u *PermissionRepository) DeletePermissionByID(permissionID int64) error {
	return u.mysqlDb.Where("id = ?", permissionID).Delete(&model.Permission{}).Error
}

//更新Permission信息
func (u *PermissionRepository) UpdatePermission(permission *model.Permission) error {
	return u.mysqlDb.Model(permission).Update(permission).Error
}

//获取结果集
func (u *PermissionRepository) FindAll() (permissionAll []model.Permission, err error) {
	return permissionAll, u.mysqlDb.Find(&permissionAll).Error
}

// 根据ID查找所有权限
func (u *PermissionRepository) FindAllPermissionById(id []int64) (permissionAll []*model.Permission, err error) {
	return permissionAll, u.mysqlDb.Find(&permissionAll, id).Error
}
