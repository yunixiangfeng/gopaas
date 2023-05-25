package repository
import (
	"github.com/jinzhu/gorm"
	"github.com/yunixiangfeng/gopaas/route/domain/model"
)
//创建需要实现的接口
type IRouteRepository interface{
    //初始化表
    InitTable() error
    //根据ID查处找数据
    FindRouteByID(int64) (*model.Route, error)
    //创建一条 route 数据
	CreateRoute(*model.Route) (int64, error)
    //根据ID删除一条 route 数据
	DeleteRouteByID(int64) error
    //修改更新数据
	UpdateRoute(*model.Route) error
    //查找route所有数据
	FindAll()([]model.Route,error)

}
//创建routeRepository
func NewRouteRepository(db *gorm.DB) IRouteRepository  {
	return &RouteRepository{mysqlDb:db}
}

type RouteRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *RouteRepository)InitTable() error  {
	return u.mysqlDb.CreateTable(&model.Route{}).Error
}

//根据ID查找Route信息
func (u *RouteRepository)FindRouteByID(routeID int64) (route *model.Route,err error) {
	route = &model.Route{}
	return route, u.mysqlDb.First(route,routeID).Error
}

//创建Route信息
func (u *RouteRepository) CreateRoute(route *model.Route) (int64, error) {
	return route.ID, u.mysqlDb.Create(route).Error
}

//根据ID删除Route信息
func (u *RouteRepository) DeleteRouteByID(routeID int64) error {
	return u.mysqlDb.Where("id = ?",routeID).Delete(&model.Route{}).Error
}

//更新Route信息
func (u *RouteRepository) UpdateRoute(route *model.Route) error {
	return u.mysqlDb.Model(route).Update(route).Error
}

//获取结果集
func (u *RouteRepository) FindAll()(routeAll []model.Route,err error) {
	return routeAll, u.mysqlDb.Find(&routeAll).Error
}

