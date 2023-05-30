package service

import (
	"github.com/yunixiangfeng/gopaas/user/domain/model"
	"github.com/yunixiangfeng/gopaas/user/domain/repository"
	"k8s.io/client-go/kubernetes"
)

//这里是接口类型
type IRoleDataService interface {
	AddRole(*model.Role) (int64, error)
	DeleteRole(int64) error
	UpdateRole(*model.Role) error
	FindRoleByID(int64) (*model.Role, error)
	FindAllRole() ([]model.Role, error)

	// 根据ID查找所有角色
	FindAllRoleByID([]int64) ([]*model.Role, error)

	// 添加权限
	AddPermission(*model.Role, []*model.Permission) error
	UpdatePermission(*model.Role, []*model.Permission) error
	DeletePermission(*model.Role, []*model.Permission) error
}

//创建
//注意：返回值 IRoleDataService 接口类型
func NewRoleDataService(roleRepository repository.IRoleRepository, clientSet *kubernetes.Clientset) IRoleDataService {
	return &RoleDataService{RoleRepository: roleRepository}
}

type RoleDataService struct {
	//注意：这里是 IRoleRepository 类型
	RoleRepository repository.IRoleRepository
}

// 根据ID查找所有角色
func (u *RoleDataService) FindAllRoleByID(id []int64) (roleAll []*model.Role, err error) {
	return u.RoleRepository.FindAllRoleById(id)
}

// 添加权限
func (u *RoleDataService) AddPermission(role *model.Role, permission []*model.Permission) error {
	return u.RoleRepository.AddPermission(role, permission)
}

func (u *RoleDataService) UpdatePermission(role *model.Role, permission []*model.Permission) error {
	return u.RoleRepository.UpdatePermission(role, permission)
}

func (u *RoleDataService) DeletePermission(role *model.Role, permission []*model.Permission) error {
	return u.RoleRepository.DeletePermission(role, permission)
}

//插入
func (u *RoleDataService) AddRole(role *model.Role) (int64, error) {
	return u.RoleRepository.CreateRole(role)
}

//删除
func (u *RoleDataService) DeleteRole(roleID int64) error {
	return u.RoleRepository.DeleteRoleByID(roleID)
}

//更新
func (u *RoleDataService) UpdateRole(role *model.Role) error {
	return u.RoleRepository.UpdateRole(role)
}

//查找
func (u *RoleDataService) FindRoleByID(roleID int64) (*model.Role, error) {
	return u.RoleRepository.FindRoleByID(roleID)
}

//查找
func (u *RoleDataService) FindAllRole() ([]model.Role, error) {
	return u.RoleRepository.FindAll()
}
