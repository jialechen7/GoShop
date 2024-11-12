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
	IComment interface {
		// GetListBackend 查询评论列表
		GetListBackend(ctx context.Context, in model.CommentGetListInput) (out *model.CommentGetListOutput, err error)
		// GetListFrontend 查询评论列表
		GetListFrontend(ctx context.Context, in model.CommentGetListInput) (out *model.CommentGetListOutput, err error)
		// UpdateBackend 更新评论
		UpdateBackend(ctx context.Context, in model.CommentUpdateInput) error
		// DeleteBackend 删除评论
		DeleteBackend(ctx context.Context, id int) error
		// AddFrontend 添加评论
		AddFrontend(ctx context.Context, in model.CommentAddInput) (out *model.CommentAddOutput, err error)
		// DeleteFrontend 删除评论
		DeleteFrontend(ctx context.Context, id int) error
	}
)

var (
	localComment IComment
)

func Comment() IComment {
	if localComment == nil {
		panic("implement not found for interface IComment, forgot register?")
	}
	return localComment
}

func RegisterComment(i IComment) {
	localComment = i
}
