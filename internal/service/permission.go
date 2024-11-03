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
	IPermission interface {
		// GetList 查询权限列表
		GetList(ctx context.Context, in model.PermissionGetListInput) (out *model.PermissionGetListOutput, err error)
		// Create 创建权限
		Create(ctx context.Context, in model.PermissionCreateInput) (out model.PermissionCreateOutput, err error)
		// Delete 删除权限
		Delete(ctx context.Context, id int) error
		// Update 修改权限
		Update(ctx context.Context, in model.PermissionUpdateInput) error
	}
)

var (
	localPermission IPermission
)

func Permission() IPermission {
	if localPermission == nil {
		panic("implement not found for interface IPermission, forgot register?")
	}
	return localPermission
}

func RegisterPermission(i IPermission) {
	localPermission = i
}
