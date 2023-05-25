package handler

import (
	"context"
    "encoding/json"
	log "github.com/asim/go-micro/v3/logger"
    route "github.com/yunixiangfeng/gopaas/route/proto/route"
	routeApi "github.com/yunixiangfeng/gopaas/routeApi/proto/routeApi"
)

type RouteApi struct{
    RouteService route.RouteService
}


// routeApi.FindRouteById 通过API向外暴露为/routeApi/findRouteById，接收http请求
// 即：/routeApi/FindRouteById 请求会调用go.micro.api.routeApi 服务的routeApi.FindRouteById 方法
func (e *RouteApi) FindRouteById(ctx context.Context, req *routeApi.Request, rsp *routeApi.Response) error {
	log.Info("Received routeApi.FindRouteById request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/routeApi/FindRouteById'}")
	rsp.Body = string(b)
	return nil
}

// routeApi.AddRoute 通过API向外暴露为/routeApi/AddRoute，接收http请求
// 即：/routeApi/AddRoute 请求会调用go.micro.api.routeApi 服务的routeApi.AddRoute 方法
func (e *RouteApi) AddRoute(ctx context.Context, req *routeApi.Request, rsp *routeApi.Response) error {
	log.Info("Received routeApi.AddRoute request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/routeApi/AddRoute'}")
	rsp.Body = string(b)
	return nil
}

// routeApi.DeleteRouteById 通过API向外暴露为/routeApi/DeleteRouteById，接收http请求
// 即：/routeApi/DeleteRouteById 请求会调用go.micro.api.routeApi 服务的 routeApi.DeleteRouteById 方法
func (e *RouteApi) DeleteRouteById(ctx context.Context, req *routeApi.Request, rsp *routeApi.Response) error {
	log.Info("Received routeApi.DeleteRouteById request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/routeApi/DeleteRouteById'}")
	rsp.Body = string(b)
	return nil
}

// routeApi.UpdateRoute 通过API向外暴露为/routeApi/UpdateRoute，接收http请求
// 即：/routeApi/UpdateRoute 请求会调用go.micro.api.routeApi 服务的routeApi.UpdateRoute 方法
func (e *RouteApi) UpdateRoute(ctx context.Context, req *routeApi.Request, rsp *routeApi.Response) error {
	log.Info("Received routeApi.UpdateRoute request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/routeApi/UpdateRoute'}")
	rsp.Body = string(b)
	return nil
}

// 默认的方法routeApi.Call 通过API向外暴露为/routeApi/call，接收http请求
// 即：/routeApi/call或/routeApi/ 请求会调用go.micro.api.routeApi 服务的routeApi.FindRouteById 方法
func (e *RouteApi) Call(ctx context.Context, req *routeApi.Request, rsp *routeApi.Response) error {
	log.Info("Received routeApi.Call request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问：Call'}")
	rsp.Body = string(b)
	return nil
}

