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
	ICoupon interface {
		// GetListBackend 查询优惠券列表
		GetListBackend(ctx context.Context, in model.CouponGetListInput) (out *model.CouponGetListOutput, err error)
		// GetListFrontend 查询可抢优惠券列表
		GetListFrontend(ctx context.Context, in model.CouponGetListAvailableInput) (out *model.CouponGetListAvailableOutput, err error)
		// Add 添加优惠券
		Add(ctx context.Context, in model.CouponAddInput) (out *model.CouponAddOutput, err error)
		// Delete 删除优惠券
		Delete(ctx context.Context, id int) error
		// Update 更新优惠券
		Update(ctx context.Context, in model.CouponUpdateInput) error
	}
)

var (
	localCoupon ICoupon
)

func Coupon() ICoupon {
	if localCoupon == nil {
		panic("implement not found for interface ICoupon, forgot register?")
	}
	return localCoupon
}

func RegisterCoupon(i ICoupon) {
	localCoupon = i
}
