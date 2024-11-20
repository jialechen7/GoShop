package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserCouponGetListCommonReq struct {
	g.Meta `path:"/user_coupon/list" tags:"用户优惠券前台" method:"get" summary:"用户优惠券列表接口"`
	CommonPaginationReq
}
type UserCouponGetListCommonRes struct {
	List  interface{} `json:"list" description:"列表"`
	Page  int         `json:"page" description:"分页码"`
	Size  int         `json:"size" description:"分页数量"`
	Total int         `json:"total" description:"数据总数"`
}

type UserCouponAddUpdateCommon struct {
	CouponId int `json:"coupon_id" form:"coupon_id" v:"required#请选择优惠券" dc:"优惠券id"`
	Status   int `json:"status" form:"status" v:"required#请选择状态" dc:"状态"`
}

type UserCouponAddReq struct {
	g.Meta `path:"/user_coupon/add" tags:"用户优惠券前台" method:"post" summary:"创建用户优惠券接口"`
	UserCouponAddUpdateCommon
}

type UserCouponAddRes struct {
	UserCouponId int `json:"user_coupon_id"`
}

type UserCouponDeleteReq struct {
	g.Meta `path:"/user_coupon/delete" method:"delete" tags:"用户优惠券前台" summary:"删除用户优惠券接口"`
	Id     int `v:"min:1#请选择需要删除的用户优惠券" dc:"用户优惠券id"`
}
type UserCouponDeleteRes struct{}

type UserCouponUpdateReq struct {
	g.Meta `path:"/user_coupon/update" method:"post" tags:"用户优惠券前台" summary:"修改用户优惠券接口"`
	Id     int `json:"id" form:"id" v:"required#请输入id" dc:"用户优惠券id"`
	UserCouponAddUpdateCommon
}
type UserCouponUpdateRes struct{}
