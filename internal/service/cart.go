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
	ICart interface {
		// GetListFrontend 查询购物车列表
		GetListFrontend(ctx context.Context, in model.CartGetListInput) (out *model.CartGetListOutput, err error)
		// AddFrontend 添加购物车
		AddFrontend(ctx context.Context, in model.CartAddInput) (out *model.CartAddOutput, err error)
		// DeleteFrontend 删除购物车
		DeleteFrontend(ctx context.Context, ids []int) error
		// UpdateFrontend 更新购物车
		UpdateFrontend(ctx context.Context, in model.CartUpdateInput) (out model.CartUpdateOutput, err error)
	}
)

var (
	localCart ICart
)

func Cart() ICart {
	if localCart == nil {
		panic("implement not found for interface ICart, forgot register?")
	}
	return localCart
}

func RegisterCart(i ICart) {
	localCart = i
}
