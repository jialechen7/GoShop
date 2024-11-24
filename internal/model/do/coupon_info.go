// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CouponInfo is the golang structure of table coupon_info for DAO operations like Where/Data.
type CouponInfo struct {
	g.Meta     `orm:"table:coupon_info, do:true"`
	Id         interface{} // 优惠券id
	Name       interface{} //
	Condition  interface{} // 满减条件 单位分
	Price      interface{} // 优惠前面值 单位分
	GoodsIds   interface{} // 可使用的goods_ids，逗号分隔
	CategoryId interface{} // 可使用的分类id
	Type       interface{} // 优惠券类型：0：普通券 1：秒杀券
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
	DeletedAt  *gtime.Time //
}
