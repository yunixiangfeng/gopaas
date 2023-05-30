package handler

import (
	"context"

	log "github.com/asim/go-micro/v3/logger"
	"github.com/yunixiangfeng/gopaas/common"
	"github.com/yunixiangfeng/gopaas/user/domain/model"
	"github.com/yunixiangfeng/gopaas/user/domain/service"
	role "github.com/yunixiangfeng/gopaas/user/proto/role"
)

type RoleHandler struct {
	//注意这里的类型是 IRoleDataService 接口类型
	RoleDataService       service.IRoleDataService
	PermissionDataService service.IPermissionDataService
}

func (e *RoleHandler) getRolePermission(rolePermission *role.RolePermission) (role *model.Role, permission []*model.Permission, err error) {
	role, err = e.RoleDataService.FindRoleByID(rolePermission.RoleId)
	if err != nil {
		common.Error(err)
		return
	}
	permission, err = e.PermissionDataService.FindAllPermissionByID(rolePermission.PermissionId)
	if err != nil {
		common.Error(err)
		return
	}
	return
}

func (e *RoleHandler) AddPermission(ctx context.Context, rolePermission *role.RolePermission, rsp *role.Response) error {
	role, permission, err := e.getRolePermission(rolePermission)
	if err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	if err := e.RoleDataService.AddPermission(role, permission); err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	return nil
}

func (e *RoleHandler) UpdatePermission(ctx context.Context, rolePermission *role.RolePermission, rsp *role.Response) error {
	role, permission, err := e.getRolePermission(rolePermission)
	if err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	if err := e.RoleDataService.UpdatePermission(role, permission); err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	return nil
}

func (e *RoleHandler) DeletePermission(ctx context.Context, rolePermission *role.RolePermission, rsp *role.Response) error {
	role, permission, err := e.getRolePermission(rolePermission)
	if err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	if err := e.RoleDataService.DeletePermission(role, permission); err != nil {
		common.Error(err)
		rsp.Msg = err.Error()
		return err
	}
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *RoleHandler) AddRole(ctx context.Context, info *role.RoleInfo, rsp *role.Response) error {
	log.Info("Received *role.AddRole request")

	return nil
}

func (e *RoleHandler) DeleteRole(ctx context.Context, req *role.RoleId, rsp *role.Response) error {
	log.Info("Received *role.DeleteRole request")

	return nil
}

func (e *RoleHandler) UpdateRole(ctx context.Context, req *role.RoleInfo, rsp *role.Response) error {
	log.Info("Received *role.UpdateRole request")

	return nil
}

func (e *RoleHandler) FindRoleByID(ctx context.Context, req *role.RoleId, rsp *role.RoleInfo) error {
	log.Info("Received *role.FindRoleByID request")

	return nil
}

func (e *RoleHandler) FindAllRole(ctx context.Context, req *role.FindAll, rsp *role.AllRole) error {
	log.Info("Received *role.FindAllRole request")

	return nil
}
