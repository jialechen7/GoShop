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
	IArticle interface {
		// GetListBackend 查询文章列表
		GetListBackend(ctx context.Context, in model.ArticleGetListInput) (out *model.ArticleGetListOutput, err error)
		// UpdateBackend 更新文章
		UpdateBackend(ctx context.Context, in model.ArticleUpdateInput) error
		// AddBackend 添加文章
		AddBackend(ctx context.Context, in model.ArticleAddInput) (out *model.ArticleAddOutput, err error)
		// DeleteBackend 删除文章
		DeleteBackend(ctx context.Context, id int) error
		// GetMyListFrontend 查询文章列表（仅用户发表的文章）
		GetMyListFrontend(ctx context.Context, in model.ArticleGetListInput) (out *model.ArticleGetListOutput, err error)
		// GetListFrontend 查询文章列表（所有人包括未登录都可查看）
		GetListFrontend(ctx context.Context, in model.ArticleGetListInput) (out *model.ArticleGetListOutput, err error)
		// AddFrontend 添加文章
		AddFrontend(ctx context.Context, in model.ArticleAddInput) (out *model.ArticleAddOutput, err error)
		// DeleteFrontend 删除文章
		DeleteFrontend(ctx context.Context, id int) error
		// DetailFrontend 查询文章详情
		DetailFrontend(ctx context.Context, id int) (out *model.ArticleDetailOutput, err error)
	}
)

var (
	localArticle IArticle
)

func Article() IArticle {
	if localArticle == nil {
		panic("implement not found for interface IArticle, forgot register?")
	}
	return localArticle
}

func RegisterArticle(i IArticle) {
	localArticle = i
}
