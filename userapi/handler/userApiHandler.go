package handler

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"

	log "github.com/asim/go-micro/v3/logger"
	"github.com/yunixiangfeng/gopaas/common"
	role "github.com/yunixiangfeng/gopaas/user/proto/role"
	user "github.com/yunixiangfeng/gopaas/user/proto/user"
	userApi "github.com/yunixiangfeng/gopaas/userApi/proto/userApi"
)

type UserApi struct {
	UserService user.UserService
	RoleService role.RoleService
}

func (e *UserApi) getPost(req *userApi.Request, key string) (string, error) {
	if _, ok := req.Post[key]; !ok {
		return "", errors.New("参数异常")
	}
	return req.Post[key].Values[0], nil
}

func (e *UserApi) getStringInt64(stringValue string) int64 {
	intValue, err := strconv.ParseInt(stringValue, 10, 64)
	if err != nil {
		common.Error(err)
		return 0
	}
	return intValue
}

func (e *UserApi) AddRole(ctx context.Context, req *userApi.Request, rsp *userApi.Response) error {
	userIdString, err := e.getPost(req, "user_id")
	if err != nil {
		common.Error(err)
		return err
	}
	userId := e.getStringInt64(userIdString)

	if _, ok := req.Post["role_id"]; !ok {
		common.Error(err)
		return err
	}
	roleId := []int64{}
	for _, v := range req.Post["role_id"].Values {
		roleId = append(roleId, e.getStringInt64(v))
	}

	rs, err := e.UserService.AddRole(ctx, &user.UserRole{
		UserId: userId,
		RoleId: roleId,
	})
	if err != nil {
		common.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(rs)
	rsp.Body = string(b)
	return nil
}

func (e *UserApi) UpdateRole(ctx context.Context, req *userApi.Request, rsp *userApi.Response) error {
	userIdString, err := e.getPost(req, "user_id")
	if err != nil {
		common.Error(err)
		return err
	}
	userId := e.getStringInt64(userIdString)

	if _, ok := req.Post["role_id"]; !ok {
		common.Error(err)
		return err
	}
	roleId := []int64{}
	for _, v := range req.Post["role_id"].Values {
		roleId = append(roleId, e.getStringInt64(v))
	}

	rs, err := e.UserService.UpdateRole(ctx, &user.UserRole{
		UserId: userId,
		RoleId: roleId,
	})
	if err != nil {
		common.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(rs)
	rsp.Body = string(b)
	return nil
}

func (e *UserApi) DeleteRole(ctx context.Context, req *userApi.Request, rsp *userApi.Response) error {
	userIdString, err := e.getPost(req, "user_id")
	if err != nil {
		common.Error(err)
		return err
	}
	userId := e.getStringInt64(userIdString)

	if _, ok := req.Post["role_id"]; !ok {
		common.Error(err)
		return err
	}
	roleId := []int64{}
	for _, v := range req.Post["role_id"].Values {
		roleId = append(roleId, e.getStringInt64(v))
	}

	rs, err := e.UserService.DeleteRole(ctx, &user.UserRole{
		UserId: userId,
		RoleId: roleId,
	})
	if err != nil {
		common.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(rs)
	rsp.Body = string(b)
	return nil
}

func (e *UserApi) IsRight(ctx context.Context, req *userApi.Request, rsp *userApi.Response) error {
	if _, ok := req.Get["user_id"]; !ok {
		return errors.New("参数异常")
	}
	idString := req.Get["user_id"].Values[0]
	userId, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		common.Error(err)
		return err
	}
	if _, ok := req.Get["user_action"]; !ok {
		return errors.New("参数异常")
	}
	action := req.Get["user_action"].Values[0]

	right, err := e.UserService.IsRight(ctx, &user.UserRight{
		UserId: userId,
		Action: action,
	})
	if err != nil {
		common.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(right)
	rsp.Body = string(b)
	return nil
}

func (e *UserApi) AddPermission(ctx context.Context, req *userApi.Request, rsp *userApi.Response) error {
	roleIdString, err := e.getPost(req, "role_id")
	if err != nil {
		common.Error(err)
		return err
	}
	roleId := e.getStringInt64(roleIdString)

	if _, ok := req.Post["permission_id"]; !ok {
		common.Error(err)
		return err
	}
	permissionId := []int64{}
	for _, v := range req.Post["permission_id"].Values {
		permissionId = append(permissionId, e.getStringInt64(v))
	}

	rs, err := e.RoleService.AddPermission(ctx, &role.RolePermission{
		PermissionId: permissionId,
		RoleId:       roleId,
	})
	if err != nil {
		common.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(rs)
	rsp.Body = string(b)
	return nil
}

func (e *UserApi) UpdatePermission(ctx context.Context, req *userApi.Request, rsp *userApi.Response) error {
	roleIdString, err := e.getPost(req, "role_id")
	if err != nil {
		common.Error(err)
		return err
	}
	roleId := e.getStringInt64(roleIdString)

	if _, ok := req.Post["permission_id"]; !ok {
		common.Error(err)
		return err
	}
	permissionId := []int64{}
	for _, v := range req.Post["permission_id"].Values {
		permissionId = append(permissionId, e.getStringInt64(v))
	}

	rs, err := e.RoleService.UpdatePermission(ctx, &role.RolePermission{
		PermissionId: permissionId,
		RoleId:       roleId,
	})
	if err != nil {
		common.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(rs)
	rsp.Body = string(b)
	return nil
}

func (e *UserApi) DeletePermission(ctx context.Context, req *userApi.Request, rsp *userApi.Response) error {
	roleIdString, err := e.getPost(req, "role_id")
	if err != nil {
		common.Error(err)
		return err
	}
	roleId := e.getStringInt64(roleIdString)

	if _, ok := req.Post["permission_id"]; !ok {
		common.Error(err)
		return err
	}
	permissionId := []int64{}
	for _, v := range req.Post["permission_id"].Values {
		permissionId = append(permissionId, e.getStringInt64(v))
	}

	rs, err := e.RoleService.DeletePermission(ctx, &role.RolePermission{
		PermissionId: permissionId,
		RoleId:       roleId,
	})
	if err != nil {
		common.Error(err)
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(rs)
	rsp.Body = string(b)
	return nil
}

// userApi.FindUserById 通过API向外暴露为/userApi/findUserById，接收http请求
// 即：/userApi/FindUserById 请求会调用go.micro.api.userApi 服务的userApi.FindUserById 方法
func (e *UserApi) FindUserById(ctx context.Context, req *userApi.Request, rsp *userApi.Response) error {
	log.Info("Received userApi.FindUserById request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/userApi/FindUserById'}")
	rsp.Body = string(b)
	return nil
}

// userApi.AddUser 通过API向外暴露为/userApi/AddUser，接收http请求
// 即：/userApi/AddUser 请求会调用go.micro.api.userApi 服务的userApi.AddUser 方法
func (e *UserApi) AddUser(ctx context.Context, req *userApi.Request, rsp *userApi.Response) error {
	log.Info("Received userApi.AddUser request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/userApi/AddUser'}")
	rsp.Body = string(b)
	return nil
}

// userApi.DeleteUserById 通过API向外暴露为/userApi/DeleteUserById，接收http请求
// 即：/userApi/DeleteUserById 请求会调用go.micro.api.userApi 服务的 userApi.DeleteUserById 方法
func (e *UserApi) DeleteUserById(ctx context.Context, req *userApi.Request, rsp *userApi.Response) error {
	log.Info("Received userApi.DeleteUserById request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/userApi/DeleteUserById'}")
	rsp.Body = string(b)
	return nil
}

// userApi.UpdateUser 通过API向外暴露为/userApi/UpdateUser，接收http请求
// 即：/userApi/UpdateUser 请求会调用go.micro.api.userApi 服务的userApi.UpdateUser 方法
func (e *UserApi) UpdateUser(ctx context.Context, req *userApi.Request, rsp *userApi.Response) error {
	log.Info("Received userApi.UpdateUser request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/userApi/UpdateUser'}")
	rsp.Body = string(b)
	return nil
}

// 默认的方法userApi.Call 通过API向外暴露为/userApi/call，接收http请求
// 即：/userApi/call或/userApi/ 请求会调用go.micro.api.userApi 服务的userApi.FindUserById 方法
func (e *UserApi) Call(ctx context.Context, req *userApi.Request, rsp *userApi.Response) error {
	log.Info("Received userApi.Call request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问：Call'}")
	rsp.Body = string(b)
	return nil
}
