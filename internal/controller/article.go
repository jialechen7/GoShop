package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/api/frontend"
	"goshop/internal/consts"
	"goshop/internal/model"
	"goshop/internal/service"

	"github.com/gogf/gf/util/gconv"
)

// Article 文章管理
var Article = cArticle{}

type cArticle struct{}

// ListBackend 查询文章列表
func (c *cArticle) ListBackend(ctx context.Context, req *backend.ArticleGetListCommonReq) (res *backend.ArticleGetListCommonRes, err error) {
	getListRes, err := service.Article().GetListBackend(ctx, model.ArticleGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.ArticleGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cArticle) UpdateBackend(ctx context.Context, req *backend.ArticleUpdateReq) (res *backend.ArticleUpdateRes, err error) {
	err = service.Article().UpdateBackend(ctx, model.ArticleUpdateInput{
		Id: req.Id,
		ArticleCreateUpdateBase: model.ArticleCreateUpdateBase{
			Title:  req.Title,
			Desc:   req.Desc,
			PicUrl: req.PicUrl,
			Detail: req.Detail,
		},
	})
	return
}

// AddBackend 添加文章
func (c *cArticle) AddBackend(ctx context.Context, req *backend.ArticleAddReq) (res *backend.ArticleAddRes, err error) {
	out, err := service.Article().AddBackend(ctx, model.ArticleAddInput{
		ArticleCreateUpdateBase: model.ArticleCreateUpdateBase{
			UserId:  gconv.Int(ctx.Value(consts.CtxAdminId)),
			Title:   req.Title,
			Desc:    req.Desc,
			PicUrl:  req.PicUrl,
			Detail:  req.Detail,
			IsAdmin: consts.ArticlePublisherAdmin,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.ArticleAddRes{
		ArticleId: out.ArticleId,
	}, nil
}

// ListMyFrontend 查询文章列表（仅用户自己）
func (c *cArticle) ListMyFrontend(ctx context.Context, req *frontend.ArticleGetMyListCommonReq) (res *frontend.ArticleGetListCommonRes, err error) {
	getListRes, err := service.Article().GetMyListFrontend(ctx, model.ArticleGetListInput{
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

// ListFrontend 查询文章列表
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

func (c *cArticle) DeleteBackend(ctx context.Context, req *backend.ArticleDeleteReq) (res *backend.ArticleDeleteRes, err error) {
	err = service.Article().DeleteBackend(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &backend.ArticleDeleteRes{}, nil
}

func (c *cArticle) AddFrontend(ctx context.Context, req *frontend.ArticleAddReq) (res *frontend.ArticleAddRes, err error) {
	out, err := service.Article().AddFrontend(ctx, model.ArticleAddInput{
		ArticleCreateUpdateBase: model.ArticleCreateUpdateBase{
			UserId:  gconv.Int(ctx.Value(consts.CtxUserId)),
			Title:   req.Title,
			Desc:    req.Desc,
			PicUrl:  req.PicUrl,
			IsAdmin: consts.ArticlePublisherFrontend,
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
		Id:         out.Id,
		UserId:     out.UserId,
		Title:      out.Title,
		Desc:       out.Desc,
		Detail:     out.Detail,
		PicUrl:     out.PicUrl,
		IsAdmin:    out.IsAdmin,
		Praise:     out.Praise,
		Collection: out.Collection,
		IsPraise:   out.IsPraise,
		IsCollect:  out.IsCollect,
		CreatedAt:  out.CreatedAt,
		UpdatedAt:  out.UpdatedAt,
	}, nil
}
