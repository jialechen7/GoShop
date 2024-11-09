package controller

import (
	"context"
	"goshop/api/frontend"
	"goshop/internal/model"
	"goshop/internal/service"
)

// Article 文章管理
var Article = cArticle{}

type cArticle struct{}

// ListFrontend 查询文章列表（仅用户自己）
func (c *cArticle) ListFrontend(ctx context.Context, req *frontend.ArticleGetListCommonReq) (res *frontend.ArticleGetListCommonRes, err error) {
	getListRes, err := service.Article().GetListFrontend(ctx, model.ArticleGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.ArticleGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cArticle) AddFrontend(ctx context.Context, req *frontend.ArticleAddReq) (res *frontend.ArticleAddRes, err error) {
	out, err := service.Article().AddFrontend(ctx, model.ArticleAddInput{
		ArticleCreateUpdateBase: model.ArticleCreateUpdateBase{
			UserId:  req.UserId,
			Title:   req.Title,
			Desc:    req.Desc,
			PicUrl:  req.PicUrl,
			IsAdmin: req.IsAdmin,
			Detail:  req.Detail,
		},
	})
	if err != nil {
		return nil, err
	}
	return &frontend.ArticleAddRes{
		ArticleId: out.ArticleId,
	}, nil
}

func (c *cArticle) DeleteFrontend(ctx context.Context, req *frontend.ArticleDeleteReq) (res *frontend.ArticleDeleteRes, err error) {
	err = service.Article().DeleteFrontend(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &frontend.ArticleDeleteRes{}, nil
}

func (c *cArticle) DetailFrontend(ctx context.Context, req *frontend.ArticleDetailReq) (res *frontend.ArticleDetailRes, err error) {
	out, err := service.Article().DetailFrontend(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &frontend.ArticleDetailRes{
		Title:     out.Title,
		Desc:      out.Desc,
		Detail:    out.Detail,
		PicUrl:    out.PicUrl,
		CreatedAt: out.CreatedAt,
	}, nil
}
