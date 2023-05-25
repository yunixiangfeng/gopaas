package handler
import (
	"context"
    "github.com/yunixiangfeng/gopaas/route/domain/service"
	log "github.com/asim/go-micro/v3/logger"
	route "github.com/yunixiangfeng/gopaas/route/proto/route"
)
type RouteHandler struct{
     //注意这里的类型是 IRouteDataService 接口类型
     RouteDataService service.IRouteDataService
}


// Call is a single request handler called via client.Call or the generated client code
func (e *RouteHandler) AddRoute(ctx context.Context,info *route.RouteInfo , rsp *route.Response) error {
	log.Info("Received *route.AddRoute request")


	return nil
}

func (e *RouteHandler) DeleteRoute(ctx context.Context, req *route.RouteId, rsp *route.Response) error {
	log.Info("Received *route.DeleteRoute request")

	return nil
}

func (e *RouteHandler) UpdateRoute(ctx context.Context, req *route.RouteInfo, rsp *route.Response) error {
	log.Info("Received *route.UpdateRoute request")

	return nil
}

func (e *RouteHandler) FindRouteByID(ctx context.Context, req *route.RouteId, rsp *route.RouteInfo) error {
	log.Info("Received *route.FindRouteByID request")

	return nil
}

func (e *RouteHandler) FindAllRoute(ctx context.Context, req *route.FindAll, rsp *route.AllRoute) error {
	log.Info("Received *route.FindAllRoute request")

	return nil
}


