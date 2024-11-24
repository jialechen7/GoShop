package backend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CouponGetListCommonReq struct {
	g.Meta `path:"/coupon/list" tags:"优惠券后台" method:"get" summary:"优惠券列表接口"`
	CommonPaginationReq
}
type CouponGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type CouponAddUpdateCommon struct {
	Name       string `json:"name" form:"name" v:"required#请输入优惠券名称" description:"优惠券名称"`
	Condition  int    `json:"condition" form:"condition" v:"required#请输入优惠券条件" description:"优惠券条件(单位分)"`
	Price      int    `json:"price" form:"price" v:"required#请输入优惠券金额" description:"优惠券金额(单位分)"`
	GoodsIds   string `json:"goods_ids" form:"goods_ids" description:"可用的商品id列表"`
	CategoryId int    `json:"category_id" form:"category_id" v:"required#请选择分类" description:"可用的分类id"`
}

type CouponAddReq struct {
	g.Meta `path:"/coupon/add" tags:"优惠券后台" method:"post" summary:"创建优惠券接口"`
	CouponAddUpdateCommon
}

type CouponAddRes struct {
	CouponId int `json:"coupon_id"`
}

type CouponDeleteReq struct {
	g.Meta `path:"/coupon/delete" method:"delete" tags:"优惠券后台" summary:"删除优惠券接口"`
	Id     int `v:"min:1#请选择需要删除的优惠券" dc:"优惠券id"`
}
type CouponDeleteRes struct{}

type CouponUpdateReq struct {
	g.Meta `path:"/coupon/update" method:"post" tags:"优惠券后台" summary:"修改优惠券接口"`
	Id     int `json:"id" form:"id" v:"required#请输入id" dc:"优惠券id"`
	CouponAddUpdateCommon
}
type CouponUpdateRes struct{}
