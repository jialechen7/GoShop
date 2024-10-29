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
	IPosition interface {
		// GetList 查询手工位列表
		GetList(ctx context.Context, in model.PositionGetListInput) (out *model.PositionGetListOutput, err error)
		// Create 创建手工位
		Create(ctx context.Context, in model.PositionCreateInput) (out model.PositionCreateOutput, err error)
		// Delete 删除手工位
		Delete(ctx context.Context, id int) error
		// Update 修改手工位
		Update(ctx context.Context, in model.PositionUpdateInput) error
	}
)

var (
	localPosition IPosition
)

func Position() IPosition {
	if localPosition == nil {
		panic("implement not found for interface IPosition, forgot register?")
	}
	return localPosition
}

func RegisterPosition(i IPosition) {
	localPosition = i
}
