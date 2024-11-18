package controller

import (
	"context"
	"goshop/api/frontend"
	"goshop/internal/model"
	"goshop/internal/service"
)

// Cart 购物车管理
var Cart = cCart{}

type cCart struct{}

// ListFrontend 查询购物车列表（仅用户自己）
func (c *cCart) ListFrontend(ctx context.Context, req *frontend.CartGetListCommonReq) (res *frontend.CartGetListCommonRes, err error) {
	getListRes, err := service.Cart().GetListFrontend(ctx, model.CartGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.CartGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cCart) AddFrontend(ctx context.Context, req *frontend.CartAddReq) (res *frontend.CartAddRes, err error) {
	out, err := service.Cart().AddFrontend(ctx, model.CartAddInput{
		CartCreateUpdateBase: model.CartCreateUpdateBase{
			GoodsOptionsId: req.GoodsOptionsId,
			Count:          req.Count,
		},
	})
	if err != nil {
		return nil, err
	}
	return &frontend.CartAddRes{
		CartId: out.CartId,
	}, nil
}

func (c *cCart) DeleteFrontend(ctx context.Context, req *frontend.CartDeleteReq) (res *frontend.CartDeleteRes, err error) {
	err = service.Cart().DeleteFrontend(ctx, req.Ids)
	if err != nil {
		return nil, err
	}
	return &frontend.CartDeleteRes{}, nil
}

// UpdateFrontend 更新购物车
func (c *cCart) UpdateFrontend(ctx context.Context, req *frontend.CartUpdateReq) (res *frontend.CartUpdateRes, err error) {
	out, err := service.Cart().UpdateFrontend(ctx, model.CartUpdateInput{
		CartCreateUpdateBase: model.CartCreateUpdateBase{
			GoodsOptionsId: req.GoodsOptionsId,
			Count:          req.Count,
		},
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.CartUpdateRes{Id: out.Id}, nil
}
