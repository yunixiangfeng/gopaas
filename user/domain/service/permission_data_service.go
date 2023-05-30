package service

import (
	"github.com/yunixiangfeng/gopaas/user/domain/model"
	"github.com/yunixiangfeng/gopaas/user/domain/repository"
	"k8s.io/client-go/kubernetes"
)

//这里是接口类型
type IPermissionDataService interface {
	AddPermission(*model.Permission) (int64, error)
	DeletePermission(int64) error
	UpdatePermission(*model.Permission) error
	FindPermissionByID(int64) (*model.Permission, error)
	FindAllPermission() ([]model.Permission, error)

	// 根据ID查询所有权限
	FindAllPermissionByID([]int64) ([]*model.Permission, error)
}

//创建
//注意：返回值 IPermissionDataService 接口类型
func NewPermissionDataService(permissionRepository repository.IPermissionRepository, clientSet *kubernetes.Clientset) IPermissionDataService {
	return &PermissionDataService{PermissionRepository: permissionRepository}
}

type PermissionDataService struct {
	//注意：这里是 IPermissionRepository 类型
	PermissionRepository repository.IPermissionRepository
}

// 根据ID查询所有权限
func (u *PermissionDataService) FindAllPermissionByID(id []int64) ([]*model.Permission, error) {
	return u.PermissionRepository.FindAllPermissionById(id)
}

//插入
func (u *PermissionDataService) AddPermission(permission *model.Permission) (int64, error) {
	return u.PermissionRepository.CreatePermission(permission)
}

//删除
func (u *PermissionDataService) DeletePermission(permissionID int64) error {
	return u.PermissionRepository.DeletePermissionByID(permissionID)
}

//更新
func (u *PermissionDataService) UpdatePermission(permission *model.Permission) error {
	return u.PermissionRepository.UpdatePermission(permission)
}

//查找
func (u *PermissionDataService) FindPermissionByID(permissionID int64) (*model.Permission, error) {
	return u.PermissionRepository.FindPermissionByID(permissionID)
}

//查找
func (u *PermissionDataService) FindAllPermission() ([]model.Permission, error) {
	return u.PermissionRepository.FindAll()
}
