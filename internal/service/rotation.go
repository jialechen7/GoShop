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
	IRotation interface {
		// GetList 查询轮播图列表
		GetList(ctx context.Context, in model.RotationGetListInput) (out *model.RotationGetListOutput, err error)
		// Create 创建轮播图
		Create(ctx context.Context, in model.RotationCreateInput) (out model.RotationCreateOutput, err error)
		// Delete 删除轮播图
		Delete(ctx context.Context, id int) error
		// Update 修改轮播图
		Update(ctx context.Context, in model.RotationUpdateInput) error
	}
)

var (
	localRotation IRotation
)

func Rotation() IRotation {
	if localRotation == nil {
		panic("implement not found for interface IRotation, forgot register?")
	}
	return localRotation
}

func RegisterRotation(i IRotation) {
	localRotation = i
}
