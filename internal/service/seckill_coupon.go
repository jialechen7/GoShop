// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goshop/internal/model"
)

type (
	ISeckillCoupon interface {
		// GetList 查询秒杀优惠券列表
		GetList(ctx context.Context, in model.SeckillCouponGetListInput) (out *model.SeckillCouponGetListOutput, err error)
		// Add 添加秒杀优惠券
		Add(ctx context.Context, in model.SeckillCouponAddInput) (out *model.SeckillCouponAddOutput, err error)
		// Delete 删除秒杀优惠券
		Delete(ctx context.Context, id int) error
		// Update 更新秒杀优惠券
		Update(ctx context.Context, in model.SeckillCouponUpdateInput) error
		// Kill 用户秒杀优惠券，注意传入参数的id是couponId
		Kill(ctx context.Context, id int) error
	}
)

var (
	localSeckillCoupon ISeckillCoupon
)

func SeckillCoupon() ISeckillCoupon {
	if localSeckillCoupon == nil {
		panic("implement not found for interface ISeckillCoupon, forgot register?")
	}
	return localSeckillCoupon
}

func RegisterSeckillCoupon(i ISeckillCoupon) {
	localSeckillCoupon = i
}
