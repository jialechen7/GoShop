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
	IGoodsOptions interface {
		// GetListBackend 查询商品规格列表
		GetListBackend(ctx context.Context, in model.GoodsOptionsGetListInput) (out *model.GoodsOptionsGetListOutput, err error)
		// UpdateBackend 更新商品规格
		UpdateBackend(ctx context.Context, in model.GoodsOptionsUpdateInput) error
		// AddBackend 添加商品规格
		AddBackend(ctx context.Context, in model.GoodsOptionsAddInput) (out *model.GoodsOptionsAddOutput, err error)
		// DeleteBackend 删除商品规格
		DeleteBackend(ctx context.Context, id int) error
		// GetListFrontend 查询商品规格列表
		GetListFrontend(ctx context.Context, in model.GoodsOptionsGetListInput) (out *model.GoodsOptionsGetListOutput, err error)
		// AddFrontend 添加商品规格
		AddFrontend(ctx context.Context, in model.GoodsOptionsAddInput) (out *model.GoodsOptionsAddOutput, err error)
		// DetailFrontend 查询商品规格详情
		DetailFrontend(ctx context.Context, id int) (out *model.GoodsOptionsDetailOutput, err error)
	}
)

var (
	localGoodsOptions IGoodsOptions
)

func GoodsOptions() IGoodsOptions {
	if localGoodsOptions == nil {
		panic("implement not found for interface IGoodsOptions, forgot register?")
	}
	return localGoodsOptions
}

func RegisterGoodsOptions(i IGoodsOptions) {
	localGoodsOptions = i
}
