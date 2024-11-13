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
	IGoods interface {
		// GetListBackend 查询商品列表
		GetListBackend(ctx context.Context, in model.GoodsGetListInput) (out *model.GoodsGetListOutput, err error)
		// UpdateBackend 更新商品
		UpdateBackend(ctx context.Context, in model.GoodsUpdateInput) error
		// AddBackend 添加商品
		AddBackend(ctx context.Context, in model.GoodsAddInput) (out *model.GoodsAddOutput, err error)
		// DeleteBackend 删除商品
		DeleteBackend(ctx context.Context, id int) error
		// GetListFrontend 查询商品列表
		GetListFrontend(ctx context.Context, in model.GoodsGetListInput) (out *model.GoodsGetListOutput, err error)
		// GetListByLevelFrontend 查询商品列表（2级分类）
		GetListByLevelFrontend(ctx context.Context, in model.GoodsGetListByLevelInput) (out *model.GoodsGetListOutput, err error)
		// AddFrontend 添加商品
		AddFrontend(ctx context.Context, in model.GoodsAddInput) (out *model.GoodsAddOutput, err error)
		// DetailFrontend 查询商品详情
		DetailFrontend(ctx context.Context, id int) (out *model.GoodsDetailOutput, err error)
		// DetailBackend 查询商品详情（管理员）
		DetailBackend(ctx context.Context, id int) (out *model.GoodsDetailOutput, err error)
	}
)

var (
	localGoods IGoods
)

func Goods() IGoods {
	if localGoods == nil {
		panic("implement not found for interface IGoods, forgot register?")
	}
	return localGoods
}

func RegisterGoods(i IGoods) {
	localGoods = i
}
