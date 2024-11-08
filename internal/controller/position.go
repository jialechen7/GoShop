package controller

import (
	"context"
	"goshop/api/backend"

	"goshop/internal/model"
	"goshop/internal/service"
)

// Position 手工位管理
var Position = cPosition{}

type cPosition struct{}

func (c *cPosition) List(ctx context.Context, req *backend.PositionGetListCommonReq) (res *backend.PositionGetListCommonRes, err error) {
	getListRes, err := service.Position().GetList(ctx, model.PositionGetListInput{
		Page: req.Page,
		Size: req.Size,
		Sort: req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &backend.PositionGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cPosition) Create(ctx context.Context, req *backend.PositionReq) (res *backend.PositionRes, err error) {
	out, err := service.Position().Create(ctx, model.PositionCreateInput{
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			GoodsName: req.GoodsName,
			GoodsId:   req.GoodsId,
			Sort:      req.Sort,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.PositionRes{PositionId: out.PositionId}, nil
}

func (c *cPosition) Delete(ctx context.Context, req *backend.PositionDeleteReq) (res *backend.PositionDeleteRes, err error) {
	err = service.Position().Delete(ctx, req.Id)
	return
}

func (c *cPosition) Update(ctx context.Context, req *backend.PositionUpdateReq) (res *backend.PositionUpdateRes, err error) {
	err = service.Position().Update(ctx, model.PositionUpdateInput{
		Id: req.Id,
		PositionCreateUpdateBase: model.PositionCreateUpdateBase{
			PicUrl:    req.PicUrl,
			Link:      req.Link,
			GoodsName: req.GoodsName,
			GoodsId:   req.GoodsId,
			Sort:      req.Sort,
		},
	})
	return
}
