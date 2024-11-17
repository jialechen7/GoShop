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

// Comment 评论管理
var Comment = cComment{}

type cComment struct{}

func (c *cComment) ListBackend(ctx context.Context, req *backend.CommentGetListCommonReq) (res *backend.CommentGetListCommonRes, err error) {
	getListRes, err := service.Comment().GetListBackend(ctx, model.CommentGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.CommentGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

// ListFrontend 查询评论列表
func (c *cComment) ListFrontend(ctx context.Context, req *frontend.CommentGetListCommonReq) (res *frontend.CommentGetListCommonRes, err error) {
	getListRes, err := service.Comment().GetListFrontend(ctx, model.CommentGetListFrontendInput{
		Page:     req.Page,
		Size:     req.Size,
		Type:     req.Type,
		ObjectId: req.ObjectId,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.CommentGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cComment) DeleteBackend(ctx context.Context, req *backend.CommentDeleteReq) (res *backend.CommentDeleteRes, err error) {
	err = service.Comment().DeleteBackend(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &backend.CommentDeleteRes{}, nil
}

func (c *cComment) AddFrontend(ctx context.Context, req *frontend.CommentAddReq) (res *frontend.CommentAddRes, err error) {
	out, err := service.Comment().AddFrontend(ctx, model.CommentAddInput{
		CommentCreateUpdateBase: model.CommentCreateUpdateBase{
			UserId:   gconv.Int(ctx.Value(consts.CtxUserId)),
			ParentId: req.ParentId,
			ObjectId: req.ObjectId,
			Type:     req.Type,
			Content:  req.Content,
		},
	})
	if err != nil {
		return nil, err
	}
	return &frontend.CommentAddRes{
		CommentId: out.CommentId,
	}, nil
}

func (c *cComment) DeleteFrontend(ctx context.Context, req *frontend.CommentDeleteReq) (res *frontend.CommentDeleteRes, err error) {
	err = service.Comment().DeleteFrontend(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &frontend.CommentDeleteRes{}, nil
}
