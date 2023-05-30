package handler

import (
	"context"
    "encoding/json"
	log "github.com/asim/go-micro/v3/logger"
    user "github.com/yunixiangfeng/gopaas/user/proto/user"
	userApi "github.com/yunixiangfeng/gopaas/userApi/proto/userApi"
)

type UserApi struct{
    UserService user.UserService
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

