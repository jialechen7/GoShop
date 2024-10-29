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
	IAdmin interface {
		// GetList 查询管理员列表
		GetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error)
		// Create 创建管理员
		Create(ctx context.Context, in model.AdminCreateInput) (out model.AdminCreateOutput, err error)
		// Delete 删除管理员
		Delete(ctx context.Context, id int) error
		// Update 修改管理员
		Update(ctx context.Context, in model.AdminUpdateInput) error
	}
)

var (
	localAdmin IAdmin
)

func Admin() IAdmin {
	if localAdmin == nil {
		panic("implement not found for interface IAdmin, forgot register?")
	}
	return localAdmin
}

func RegisterAdmin(i IAdmin) {
	localAdmin = i
}
