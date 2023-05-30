package handler

import (
	"context"

	log "github.com/asim/go-micro/v3/logger"
	"github.com/yunixiangfeng/gopaas/user/domain/service"
	permission "github.com/yunixiangfeng/gopaas/user/proto/permission"
)

type PermissionHandler struct {
	//注意这里的类型是 IPermissionDataService 接口类型
	PermissionDataService service.IPermissionDataService
}

// Call is a single request handler called via client.Call or the generated client code
func (e *PermissionHandler) AddPermission(ctx context.Context, info *permission.PermissionInfo, rsp *permission.Response) error {
	log.Info("Received *permission.AddPermission request")

	return nil
}

func (e *PermissionHandler) DeletePermission(ctx context.Context, req *permission.PermissionId, rsp *permission.Response) error {
	log.Info("Received *permission.DeletePermission request")

	return nil
}

func (e *PermissionHandler) UpdatePermission(ctx context.Context, req *permission.PermissionInfo, rsp *permission.Response) error {
	log.Info("Received *permission.UpdatePermission request")

	return nil
}

func (e *PermissionHandler) FindPermissionByID(ctx context.Context, req *permission.PermissionId, rsp *permission.PermissionInfo) error {
	log.Info("Received *permission.FindPermissionByID request")

	return nil
}

func (e *PermissionHandler) FindAllPermission(ctx context.Context, req *permission.FindAll, rsp *permission.AllPermission) error {
	log.Info("Received *permission.FindAllPermission request")

	return nil
}
