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
		// GetListFrontend 查询文章列表（仅用户发表的文章）
		GetListFrontend(ctx context.Context, in model.ArticleGetListInput) (out *model.ArticleGetListOutput, err error)
		// AddFrontend 添加文章
		AddFrontend(ctx context.Context, in model.ArticleAddInput) (out *model.ArticleAddOutput, err error)
		// DeleteFrontend 删除文章
		DeleteFrontend(ctx context.Context, id int) error
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
