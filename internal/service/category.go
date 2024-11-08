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
	ICategory interface {
		// GetList 查询分类列表
		GetList(ctx context.Context, in model.CategoryGetListInput) (out *model.CategoryGetListOutput, err error)
		// GetListFrontend 查询分类列表（仅用户自己的分类）
		GetListFrontend(ctx context.Context, in model.CategoryGetListWithParentIdInput) (out *model.CategoryGetListOutput, err error)
		// Add 添加分类
		Add(ctx context.Context, in model.CategoryAddInput) (out model.CategoryAddOutput, err error)
		// Delete 删除分类
		Delete(ctx context.Context, id int) error
	}
)

var (
	localCategory ICategory
)

func Category() ICategory {
	if localCategory == nil {
		panic("implement not found for interface ICategory, forgot register?")
	}
	return localCategory
}

func RegisterCategory(i ICategory) {
	localCategory = i
}
