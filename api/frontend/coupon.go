package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CouponGetListCommonReq struct {
	g.Meta  `path:"/coupon/list" tags:"优惠券前台" method:"get" summary:"优惠券列表接口"`
	GoodsId int `json:"goods_id"`
}
type CouponGetListCommonRes struct {
	List interface{} `json:"list" description:"列表"`
}

type CouponGetListAvailableReq struct {
	g.Meta          `path:"/coupon/available" tags:"优惠券前台" method:"post" summary:"可用优惠券列表接口"`
	OrderGoodsInfos []*OrderGoodsInfo `json:"order_goods_infos" v:"required#请选择订单商品信息"`
}

type CouponGetListAvailableRes struct {
	AvailableList   interface{} `json:"available_list" description:"可用列表"`
	UnavailableList interface{} `json:"unavailable_list" description:"不可用列表"`
}

type OrderGoodsInfo struct {
	GoodsId        int `json:"goods_id"`
	GoodsOptionsId int `json:"goods_options_id"`
	Count          int `json:"count"`
	Price          int `json:"price"`
}
