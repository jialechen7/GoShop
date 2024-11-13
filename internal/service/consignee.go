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
	IConsignee interface {
		// GetListFrontend 查询收货人列表
		GetListFrontend(ctx context.Context, in model.ConsigneeGetListInput) (out *model.ConsigneeGetListOutput, err error)
		// GetListBackend 查询文章列表
		GetListBackend(ctx context.Context, in model.ConsigneeGetListBackendInput) (out *model.ConsigneeGetListOutput, err error)
		// UpdateBackend 更新收货人
		UpdateBackend(ctx context.Context, in model.ConsigneeUpdateInput) error
		// AddBackend 添加收货人
		AddBackend(ctx context.Context, in model.ConsigneeAddInput) (out *model.ConsigneeAddOutput, err error)
		// DeleteBackend 删除收货人
		DeleteBackend(ctx context.Context, id int) error
		// AddFrontend 添加收货人
		AddFrontend(ctx context.Context, in model.ConsigneeAddInput) (out *model.ConsigneeAddOutput, err error)
		// DeleteFrontend 删除收货人
		DeleteFrontend(ctx context.Context, id int) error
		// UpdateFrontend 更新收货人
		UpdateFrontend(ctx context.Context, in model.ConsigneeUpdateInput) error
		UnsetDefault(ctx context.Context) error
	}
)

var (
	localConsignee IConsignee
)

func Consignee() IConsignee {
	if localConsignee == nil {
		panic("implement not found for interface IConsignee, forgot register?")
	}
	return localConsignee
}

func RegisterConsignee(i IConsignee) {
	localConsignee = i
}
