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
