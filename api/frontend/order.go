package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type OrderGetListCommonReq struct {
	g.Meta `path:"/order/list" tags:"Order" method:"get" summary:"订单列表接口"`
	Status int `json:"status" in:"query" v:"required#请选择订单状态"`
	CommonPaginationReq
}
type OrderGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type OrderAddReq struct {
	g.Meta `path:"/order/add" tags:"Order" method:"post" summary:"创建订单接口"`
}

type OrderAddRes struct {
	OrderId int `json:"orderId"`
}

type OrderDeleteReq struct {
	g.Meta `path:"/order/delete" method:"delete" tags:"Order" summary:"删除订单接口"`
	Id     int `v:"min:1#请选择需要删除的订单" dc:"订单id"`
}
type OrderDeleteRes struct{}

type OrderUpdateReq struct {
	g.Meta `path:"/order/update" method:"post" tags:"订单" summary:"修改订单接口"`
}
type OrderUpdateRes struct{}
