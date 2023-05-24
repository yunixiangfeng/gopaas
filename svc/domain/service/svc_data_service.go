package service

import (
	"github.com/yunixiangfeng/gopaas/svc/domain/model"
	"github.com/yunixiangfeng/gopaas/svc/domain/repository"
	"k8s.io/client-go/kubernetes"
)

//这里是接口类型
type ISvcDataService interface {
	AddSvc(*model.Svc) (int64, error)
	DeleteSvc(int64) error
	UpdateSvc(*model.Svc) error
	FindSvcByID(int64) (*model.Svc, error)
	FindAllSvc() ([]model.Svc, error)
}

//创建
//注意：返回值 ISvcDataService 接口类型
func NewSvcDataService(svcRepository repository.ISvcRepository, clientSet *kubernetes.Clientset) ISvcDataService {
	return &SvcDataService{SvcRepository: svcRepository, K8sClientSet: clientSet}
}

type SvcDataService struct {
	//注意：这里是 ISvcRepository 类型
	SvcRepository repository.ISvcRepository
	K8sClientSet  *kubernetes.Clientset
}

//插入
func (u *SvcDataService) AddSvc(svc *model.Svc) (int64, error) {
	return u.SvcRepository.CreateSvc(svc)
}

//删除
func (u *SvcDataService) DeleteSvc(svcID int64) error {
	return u.SvcRepository.DeleteSvcByID(svcID)
}

//更新
func (u *SvcDataService) UpdateSvc(svc *model.Svc) error {
	return u.SvcRepository.UpdateSvc(svc)
}

//查找
func (u *SvcDataService) FindSvcByID(svcID int64) (*model.Svc, error) {
	return u.SvcRepository.FindSvcByID(svcID)
}

//查找
func (u *SvcDataService) FindAllSvc() ([]model.Svc, error) {
	return u.SvcRepository.FindAll()
}
