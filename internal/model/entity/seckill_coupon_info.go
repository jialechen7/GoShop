// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SeckillCouponInfo is the golang structure for table seckill_coupon_info.
type SeckillCouponInfo struct {
	Id        int         `json:"id"        orm:"id"         description:"秒杀优惠券id"`
	CouponId  int         `json:"couponId"  orm:"coupon_id"  description:"优惠券id"`
	Stock     int         `json:"stock"     orm:"stock"      description:"库存"`
	StartTime *gtime.Time `json:"startTime" orm:"start_time" description:"开始时间"`
	EndTime   *gtime.Time `json:"endTime"   orm:"end_time"   description:"结束时间"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`
}
