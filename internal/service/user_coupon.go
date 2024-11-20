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
	IUserCoupon interface {
		// GetList 查询优惠券列表
		GetList(ctx context.Context, in model.UserCouponGetListInput) (out *model.UserCouponGetListOutput, err error)
		// Add 添加优惠券
		Add(ctx context.Context, in model.UserCouponAddInput) (out *model.UserCouponAddOutput, err error)
		// Delete 删除优惠券
		Delete(ctx context.Context, id int) error
		// Update 更新优惠券
		Update(ctx context.Context, in model.UserCouponUpdateInput) error
	}
)

var (
	localUserCoupon IUserCoupon
)

func UserCoupon() IUserCoupon {
	if localUserCoupon == nil {
		panic("implement not found for interface IUserCoupon, forgot register?")
	}
	return localUserCoupon
}

func RegisterUserCoupon(i IUserCoupon) {
	localUserCoupon = i
}
