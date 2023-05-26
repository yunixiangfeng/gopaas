package service

import (
	"github.com/yunixiangfeng/gopaas/volume/domain/model"
	"github.com/yunixiangfeng/gopaas/volume/domain/repository"
	"k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
)

//这里是接口类型
type IVolumeDataService interface {
	AddVolume(*model.Volume) (int64 , error)
	DeleteVolume(int64) error
	UpdateVolume(*model.Volume) error
	FindVolumeByID(int64) (*model.Volume, error)
	FindAllVolume() ([]model.Volume, error)
}


//创建
//注意：返回值 IVolumeDataService 接口类型
func NewVolumeDataService(volumeRepository repository.IVolumeRepository,clientSet *kubernetes.Clientset) IVolumeDataService{
	return &VolumeDataService{ VolumeRepository:volumeRepository, K8sClientSet: clientSet,deployment:&v1.Deployment{}}
}

type VolumeDataService struct {
    //注意：这里是 IVolumeRepository 类型
	VolumeRepository repository.IVolumeRepository
	K8sClientSet  *kubernetes.Clientset
	deployment  *v1.Deployment
}


//插入
func (u *VolumeDataService) AddVolume(volume *model.Volume) (int64 ,error) {
	 return u.VolumeRepository.CreateVolume(volume)
}

//删除
func (u *VolumeDataService) DeleteVolume(volumeID int64) error {
	return u.VolumeRepository.DeleteVolumeByID(volumeID)
}

//更新
func (u *VolumeDataService) UpdateVolume(volume *model.Volume) error {
	return u.VolumeRepository.UpdateVolume(volume)
}

//查找
func (u *VolumeDataService) FindVolumeByID(volumeID int64) (*model.Volume, error) {
	return u.VolumeRepository.FindVolumeByID(volumeID)
}

//查找
func (u *VolumeDataService) FindAllVolume() ([]model.Volume, error) {
	return u.VolumeRepository.FindAll()
}

