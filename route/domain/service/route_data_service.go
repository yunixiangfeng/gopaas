package service

import (
	"github.com/yunixiangfeng/gopaas/route/domain/model"
	"github.com/yunixiangfeng/gopaas/route/domain/repository"
	"k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
)

//这里是接口类型
type IRouteDataService interface {
	AddRoute(*model.Route) (int64 , error)
	DeleteRoute(int64) error
	UpdateRoute(*model.Route) error
	FindRouteByID(int64) (*model.Route, error)
	FindAllRoute() ([]model.Route, error)
}


//创建
//注意：返回值 IRouteDataService 接口类型
func NewRouteDataService(routeRepository repository.IRouteRepository,clientSet *kubernetes.Clientset) IRouteDataService{
	return &RouteDataService{ RouteRepository:routeRepository, K8sClientSet: clientSet,deployment:&v1.Deployment{}}
}

type RouteDataService struct {
    //注意：这里是 IRouteRepository 类型
	RouteRepository repository.IRouteRepository
	K8sClientSet  *kubernetes.Clientset
	deployment  *v1.Deployment
}


//插入
func (u *RouteDataService) AddRoute(route *model.Route) (int64 ,error) {
	 return u.RouteRepository.CreateRoute(route)
}

//删除
func (u *RouteDataService) DeleteRoute(routeID int64) error {
	return u.RouteRepository.DeleteRouteByID(routeID)
}

//更新
func (u *RouteDataService) UpdateRoute(route *model.Route) error {
	return u.RouteRepository.UpdateRoute(route)
}

//查找
func (u *RouteDataService) FindRouteByID(routeID int64) (*model.Route, error) {
	return u.RouteRepository.FindRouteByID(routeID)
}

//查找
func (u *RouteDataService) FindAllRoute() ([]model.Route, error) {
	return u.RouteRepository.FindAll()
}

