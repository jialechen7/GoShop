package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type OrderGetListCommonReq struct {
	g.Meta `path:"/order/list" tags:"订单前台" method:"get" summary:"订单列表接口"`
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
	g.Meta           `path:"/order/add" tags:"订单前台" method:"post" summary:"创建订单接口"`
	PayType          int                  `json:"pay_type" v:"required|in:1,2,3#请选择支付方式"`
	Remark           string               `json:"remark"`
	Status           int                  `json:"status" v:"required#请选择订单状态"`
	Price            int                  `json:"price" v:"required#请输入订单价格"`
	ConsigneeName    string               `json:"consignee_name" v:"required#请输入收货人姓名"`
	ConsigneePhone   string               `json:"consignee_phone" v:"required#请输入收货人电话"`
	ConsigneeAddress string               `json:"consignee_address" v:"required#请输入收货人地址"`
	OrderGoodsInfos  []*OrderAddGoodsInfo `json:"order_goods_infos" v:"required#请选择订单商品信息"`
}

type OrderAddGoodsInfo struct {
	GoodsId        int    `json:"goods_id"`
	GoodsOptionsId int    `json:"goods_options_id"`
	Count          int    `json:"count"`
	Remark         string `json:"remark"`
	Price          int    `json:"price"`
	CouponPrice    int    `json:"coupon_price"`
	ActualPrice    int    `json:"actual_price"`
}

type OrderAddRes struct {
	OrderId int `json:"order_id"`
}

type OrderDeleteReq struct {
	g.Meta `path:"/order/delete" method:"delete" tags:"订单前台" summary:"删除订单接口"`
	Id     int `v:"min:1#请选择需要删除的订单" dc:"订单id"`
}
type OrderDeleteRes struct{}

type OrderUpdateReq struct {
	g.Meta `path:"/order/update" method:"post" tags:"订单前台" summary:"修改订单接口"`
}
type OrderUpdateRes struct{}
