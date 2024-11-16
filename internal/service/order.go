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
	IOrder interface {
		// GetList 查询订单列表
		GetList(ctx context.Context, in model.OrderGetListInput) (out *model.OrderGetListOutput, err error)
		// GetListFrontend 查询订单列表（仅用户自己的订单）
		GetListFrontend(ctx context.Context, in model.OrderGetListWithStatusInput) (out *model.OrderGetListOutput, err error)
		AddFrontend(ctx context.Context, in model.OrderAddInput) (out *model.OrderAddOutput, err error)
	}
)

var (
	localOrder IOrder
)

func Order() IOrder {
	if localOrder == nil {
		panic("implement not found for interface IOrder, forgot register?")
	}
	return localOrder
}

func RegisterOrder(i IOrder) {
	localOrder = i
}
