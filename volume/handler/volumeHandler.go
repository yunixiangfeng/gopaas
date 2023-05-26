package handler
import (
	"context"
    "github.com/yunixiangfeng/gopaas/volume/domain/service"
	log "github.com/asim/go-micro/v3/logger"
	volume "github.com/yunixiangfeng/gopaas/volume/proto/volume"
)
type VolumeHandler struct{
     //注意这里的类型是 IVolumeDataService 接口类型
     VolumeDataService service.IVolumeDataService
}


// Call is a single request handler called via client.Call or the generated client code
func (e *VolumeHandler) AddVolume(ctx context.Context,info *volume.VolumeInfo , rsp *volume.Response) error {
	log.Info("Received *volume.AddVolume request")


	return nil
}

func (e *VolumeHandler) DeleteVolume(ctx context.Context, req *volume.VolumeId, rsp *volume.Response) error {
	log.Info("Received *volume.DeleteVolume request")

	return nil
}

func (e *VolumeHandler) UpdateVolume(ctx context.Context, req *volume.VolumeInfo, rsp *volume.Response) error {
	log.Info("Received *volume.UpdateVolume request")

	return nil
}

func (e *VolumeHandler) FindVolumeByID(ctx context.Context, req *volume.VolumeId, rsp *volume.VolumeInfo) error {
	log.Info("Received *volume.FindVolumeByID request")

	return nil
}

func (e *VolumeHandler) FindAllVolume(ctx context.Context, req *volume.FindAll, rsp *volume.AllVolume) error {
	log.Info("Received *volume.FindAllVolume request")

	return nil
}


