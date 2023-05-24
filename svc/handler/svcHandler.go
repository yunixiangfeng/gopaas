package handler
import (
	"context"
    "github.com/yunixiangfeng/gopaas/svc/domain/service"
	log "github.com/asim/go-micro/v3/logger"
	svc "github.com/yunixiangfeng/gopaas/svc/proto/svc"
)
type SvcHandler struct{
     //注意这里的类型是 ISvcDataService 接口类型
     SvcDataService service.ISvcDataService
}


// Call is a single request handler called via client.Call or the generated client code
func (e *SvcHandler) AddSvc(ctx context.Context,info *svc.SvcInfo , rsp *svc.Response) error {
	log.Info("Received *svc.AddSvc request")


	return nil
}

func (e *SvcHandler) DeleteSvc(ctx context.Context, req *svc.SvcId, rsp *svc.Response) error {
	log.Info("Received *svc.DeleteSvc request")

	return nil
}

func (e *SvcHandler) UpdateSvc(ctx context.Context, req *svc.SvcInfo, rsp *svc.Response) error {
	log.Info("Received *svc.UpdateSvc request")

	return nil
}

func (e *SvcHandler) FindSvcByID(ctx context.Context, req *svc.SvcId, rsp *svc.SvcInfo) error {
	log.Info("Received *svc.FindSvcByID request")

	return nil
}

func (e *SvcHandler) FindAllSvc(ctx context.Context, req *svc.FindAll, rsp *svc.AllSvc) error {
	log.Info("Received *svc.FindAllSvc request")

	return nil
}


