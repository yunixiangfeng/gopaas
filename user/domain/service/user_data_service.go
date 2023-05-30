package service

import (
	"github.com/yunixiangfeng/gopaas/user/domain/model"
	"github.com/yunixiangfeng/gopaas/user/domain/repository"
	"k8s.io/client-go/kubernetes"
)

//这里是接口类型
type IUserDataService interface {
	AddUser(*model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(*model.User) error
	FindUserByID(int64) (*model.User, error)
	FindAllUser() ([]model.User, error)

	// 分配角色
	AddRole(*model.User, []*model.Role) error
	UpdateRole(*model.User, []*model.Role) error
	DeleteRole(*model.User, []*model.Role) error
	// 判断用户是否有权限
	IsRight(string, int64) bool
}

//创建
//注意：返回值 IUserDataService 接口类型
func NewUserDataService(userRepository repository.IUserRepository, clientSet *kubernetes.Clientset) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

type UserDataService struct {
	//注意：这里是 IUserRepository 类型
	UserRepository repository.IUserRepository
}

// 分配角色
func (u *UserDataService) AddRole(user *model.User, role []*model.Role) error {
	return u.UserRepository.AddRole(user, role)
}
func (u *UserDataService) UpdateRole(user *model.User, role []*model.Role) error {
	return u.UserRepository.UpdateRole(user, role)
}

func (u *UserDataService) DeleteRole(user *model.User, role []*model.Role) error {
	return u.UserRepository.DeleteRole(user, role)
}

// 判断用户是否有权限
func (u *UserDataService) IsRight(action string, userID int64) bool {
	return u.UserRepository.IsRight(action, userID)
}

//插入
func (u *UserDataService) AddUser(user *model.User) (int64, error) {
	return u.UserRepository.CreateUser(user)
}

//删除
func (u *UserDataService) DeleteUser(userID int64) error {
	return u.UserRepository.DeleteUserByID(userID)
}

//更新
func (u *UserDataService) UpdateUser(user *model.User) error {
	return u.UserRepository.UpdateUser(user)
}

//查找
func (u *UserDataService) FindUserByID(userID int64) (*model.User, error) {
	return u.UserRepository.FindUserByID(userID)
}

//查找
func (u *UserDataService) FindAllUser() ([]model.User, error) {
	return u.UserRepository.FindAll()
}
