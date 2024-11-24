// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SeckillCouponInfo is the golang structure of table seckill_coupon_info for DAO operations like Where/Data.
type SeckillCouponInfo struct {
	g.Meta    `orm:"table:seckill_coupon_info, do:true"`
	Id        interface{} // 秒杀优惠券id
	CouponId  interface{} // 优惠券id
	Stock     interface{} // 库存
	StartTime *gtime.Time // 开始时间
	EndTime   *gtime.Time // 结束时间
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
