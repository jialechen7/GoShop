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
	IPraise interface {
		// GetListFrontend 查询点赞列表（仅用户发表的点赞）
		GetListFrontend(ctx context.Context, in model.PraiseGetListInput) (out *model.PraiseGetListOutput, err error)
		// AddFrontend 添加点赞
		AddFrontend(ctx context.Context, in model.PraiseAddInput) (out *model.PraiseAddOutput, err error)
		// DeleteFrontend 删除点赞
		DeleteFrontend(ctx context.Context, id int) error
	}
)

var (
	localPraise IPraise
)

func Praise() IPraise {
	if localPraise == nil {
		panic("implement not found for interface IPraise, forgot register?")
	}
	return localPraise
}

func RegisterPraise(i IPraise) {
	localPraise = i
}
