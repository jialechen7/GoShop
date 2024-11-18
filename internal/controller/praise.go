package controller

import (
	"context"
	"goshop/api/frontend"
	"goshop/internal/model"
	"goshop/internal/service"
)

// Praise 点赞管理
var Praise = cPraise{}

type cPraise struct{}

// ListFrontend 查询点赞列表（仅用户自己）
func (c *cPraise) ListFrontend(ctx context.Context, req *frontend.PraiseGetListCommonReq) (res *frontend.PraiseGetListCommonRes, err error) {
	getListRes, err := service.Praise().GetListFrontend(ctx, model.PraiseGetListInput{
		Page: req.Page,
		Size: req.Size,
		Type: req.Type,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.PraiseGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cPraise) AddFrontend(ctx context.Context, req *frontend.PraiseAddReq) (res *frontend.PraiseAddRes, err error) {
	out, err := service.Praise().AddFrontend(ctx, model.PraiseAddInput{
		PraiseCreateUpdateBase: model.PraiseCreateUpdateBase{
			Type:     req.Type,
			ObjectId: req.ObjectId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &frontend.PraiseAddRes{
		PraiseId: out.PraiseId,
	}, nil
}

func (c *cPraise) DeleteFrontend(ctx context.Context, req *frontend.PraiseDeleteReq) (res *frontend.PraiseDeleteRes, err error) {
	err = service.Praise().DeleteFrontend(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &frontend.PraiseDeleteRes{}, nil
}

func (c *cPraise) DeleteByTypeFrontend(ctx context.Context, req *frontend.PraiseDeleteByTypeReq) (res *frontend.PraiseDeleteByTypeRes, err error) {
	err = service.Praise().DeleteByTypeFrontend(ctx, model.PraiseDeleteByTypeInput{
		Type:     req.Type,
		ObjectId: req.ObjectId,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.PraiseDeleteByTypeRes{}, nil
}
