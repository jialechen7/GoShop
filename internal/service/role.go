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
	IRole interface {
		// Create 添加角色
		Create(ctx context.Context, in model.RoleCreateInput) (out model.RoleCreateOutput, err error)
		// AddPermission 添加角色权限
		AddPermission(ctx context.Context, in model.RoleAddPermissionInput) (out model.RoleAddPermissionOutput, err error)
		// DeletePermission 删除角色权限
		DeletePermission(ctx context.Context, in model.RoleDeletePermissionInput) error
		// Delete 删除角色
		Delete(ctx context.Context, id int) error
		// Update 修改角色
		Update(ctx context.Context, in model.RoleUpdateInput) error
		// GetList 查询角色列表
		GetList(ctx context.Context, in model.RoleGetListInput) (out *model.RoleGetListOutput, err error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
