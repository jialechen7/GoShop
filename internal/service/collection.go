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
	ICollection interface {
		// GetListFrontend 查询收藏列表（仅用户发表的收藏）
		GetListFrontend(ctx context.Context, in model.CollectionGetListInput) (out *model.CollectionGetListOutput, err error)
		// AddFrontend 添加收藏
		AddFrontend(ctx context.Context, in model.CollectionAddInput) (out *model.CollectionAddOutput, err error)
		// DeleteFrontend 删除收藏
		DeleteFrontend(ctx context.Context, id int) error
		// DeleteByTypeFrontend 删除收藏（根据类型）
		DeleteByTypeFrontend(ctx context.Context, in model.CollectionDeleteByTypeInput) error
	}
)

var (
	localCollection ICollection
)

func Collection() ICollection {
	if localCollection == nil {
		panic("implement not found for interface ICollection, forgot register?")
	}
	return localCollection
}

func RegisterCollection(i ICollection) {
	localCollection = i
}
