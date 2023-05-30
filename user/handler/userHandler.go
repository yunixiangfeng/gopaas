package handler

import (
	"context"

	log "github.com/asim/go-micro/v3/logger"
	"github.com/yunixiangfeng/gopaas/common"
	"github.com/yunixiangfeng/gopaas/user/domain/model"
	"github.com/yunixiangfeng/gopaas/user/domain/service"
	user "github.com/yunixiangfeng/gopaas/user/proto/user"
)

type UserHandler struct {
	//注意这里的类型是 IUserDataService 接口类型
	UserDataService service.IUserDataService
	RoleDataService service.IRoleDataService
}

func (e *UserHandler) getUserRole(userRole *user.UserRole) (user *model.User, role []*model.Role, err error) {
	user, err = e.UserDataService.FindUserByID(userRole.UserId)
	if err != nil {
		common.Error(err)
		return
	}
	role, err = e.RoleDataService.FindAllRoleByID(userRole.RoleId)
	if err != nil {
		common.Error(err)
		return
	}
	return
}

func (e *UserHandler) AddRole(ctx context.Context, userRole *user.UserRole, rsp *user.Response) error {
	// 查找用户和对应的角色
	user, role, err := e.getUserRole(userRole)
	if err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	// 添加对应的角色
	if err := e.UserDataService.AddRole(user, role); err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	return nil
}

func (e *UserHandler) UpdateRole(ctx context.Context, userRole *user.UserRole, rsp *user.Response) error {
	// 查找用户和对应的角色
	user, role, err := e.getUserRole(userRole)
	if err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	// 更新对应的角色
	if err := e.UserDataService.UpdateRole(user, role); err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	return nil
}

func (e *UserHandler) DeleteRole(ctx context.Context, userRole *user.UserRole, rsp *user.Response) error {
	// 查找用户和对应的角色
	user, role, err := e.getUserRole(userRole)
	if err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	// 删除对应的角色
	if err := e.UserDataService.DeleteRole(user, role); err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	return nil
}

func (e *UserHandler) IsRight(ctx context.Context, userRight *user.UserRight, rg *user.Right) error {
	rg.Access = e.UserDataService.IsRight(userRight.Action, userRight.UserId)
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *UserHandler) AddUser(ctx context.Context, info *user.UserInfo, rsp *user.Response) error {
	log.Info("Received *user.AddUser request")

	return nil
}

func (e *UserHandler) DeleteUser(ctx context.Context, req *user.UserId, rsp *user.Response) error {
	log.Info("Received *user.DeleteUser request")

	return nil
}

func (e *UserHandler) UpdateUser(ctx context.Context, req *user.UserInfo, rsp *user.Response) error {
	log.Info("Received *user.UpdateUser request")

	return nil
}

func (e *UserHandler) FindUserByID(ctx context.Context, req *user.UserId, rsp *user.UserInfo) error {
	log.Info("Received *user.FindUserByID request")

	return nil
}

func (e *UserHandler) FindAllUser(ctx context.Context, req *user.FindAll, rsp *user.AllUser) error {
	log.Info("Received *user.FindAllUser request")

	return nil
}
