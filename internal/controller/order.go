package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/internal/model"
	"goshop/internal/service"
)

// Order 订单管理
var Order = cOrder{}

type cOrder struct{}

func (a *cOrder) List(ctx context.Context, req *backend.OrderGetListCommonReq) (res *backend.OrderGetListCommonRes, err error) {
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
