package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/api/frontend"
	"goshop/internal/consts"
	"goshop/internal/model"
	"goshop/internal/service"
	"goshop/utility"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// Order 订单管理
var Order = cOrder{}

type cOrder struct{}

// List 查询订单列表
func (c *cOrder) List(ctx context.Context, req *backend.OrderGetListCommonReq) (res *backend.OrderGetListCommonRes, err error) {
	getListRes, err := service.Order().GetList(ctx, model.OrderGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.OrderGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

// ListFrontend 查询订单列表（仅用户自己的订单）
func (c *cOrder) ListFrontend(ctx context.Context, req *frontend.OrderGetListCommonReq) (res *frontend.OrderGetListCommonRes, err error) {
	getListRes, err := service.Order().GetListFrontend(ctx, model.OrderGetListWithStatusInput{
		Page:   req.Page,
		Size:   req.Size,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.OrderGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cOrder) AddFrontend(ctx context.Context, req *frontend.OrderAddReq) (res *frontend.OrderAddRes, err error) {
	out, err := service.Order().AddFrontend(ctx, model.OrderAddInput{
		OrderAddUpdateBase: model.OrderAddUpdateBase{
			Number:           utility.GetOrderNum(),
			UserId:           gconv.Int(ctx.Value(consts.CtxUserId)),
			PayType:          req.PayType,
			Remark:           req.Remark,
			PayAt:            gtime.Now(),
			Status:           req.Status,
			ConsigneeName:    req.ConsigneeName,
			ConsigneePhone:   req.ConsigneePhone,
			ConsigneeAddress: req.ConsigneeAddress,
			Price:            req.Price,
		},
	})
	if err != nil {
		return nil, err
	}
	return &frontend.OrderAddRes{
		OrderId: out.OrderId,
	}, nil
}
