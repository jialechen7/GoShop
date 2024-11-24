package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/api/frontend"
	"goshop/internal/model"
	"goshop/internal/service"
)

// SeckillCoupon 秒杀优惠券管理
var SeckillCoupon = cSeckillCoupon{}

type cSeckillCoupon struct{}

// List 查询秒杀优惠券列表
func (c *cSeckillCoupon) List(ctx context.Context, req *backend.SeckillCouponGetListCommonReq) (res *backend.SeckillCouponGetListCommonRes, err error) {
	getListRes, err := service.SeckillCoupon().GetList(ctx, model.SeckillCouponGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.SeckillCouponGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cSeckillCoupon) Add(ctx context.Context, req *backend.SeckillCouponAddReq) (res *backend.SeckillCouponAddRes, err error) {
	out, err := service.SeckillCoupon().Add(ctx, model.SeckillCouponAddInput{
		Name:       req.Name,
		Condition:  req.Condition,
		Price:      req.Price,
		GoodsIds:   req.GoodsIds,
		CategoryId: req.CategoryId,
		SeckillCouponCreateUpdateBase: model.SeckillCouponCreateUpdateBase{
			Stock:     req.Stock,
			StartTime: req.StartTime,
			EndTime:   req.EndTime,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.SeckillCouponAddRes{
		SeckillCouponId: out.SeckillCouponId,
	}, nil
}

func (c *cSeckillCoupon) Delete(ctx context.Context, req *backend.SeckillCouponDeleteReq) (res *backend.SeckillCouponDeleteRes, err error) {
	err = service.SeckillCoupon().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &backend.SeckillCouponDeleteRes{}, nil
}

func (c *cSeckillCoupon) Update(ctx context.Context, req *backend.SeckillCouponUpdateReq) (res *backend.SeckillCouponUpdateRes, err error) {
	err = service.SeckillCoupon().Update(ctx, model.SeckillCouponUpdateInput{
		Id: req.Id,
		SeckillCouponCreateUpdateBase: model.SeckillCouponCreateUpdateBase{
			Stock:     req.Stock,
			StartTime: req.StartTime,
			EndTime:   req.EndTime,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.SeckillCouponUpdateRes{}, nil
}

// Kill 用户秒杀优惠券
func (c *cSeckillCoupon) Kill(ctx context.Context, req *frontend.SeckillCouponKillReq) (res *frontend.SeckillCouponKillRes, err error) {
	err = service.SeckillCoupon().Kill(ctx, req.CouponId)
	if err != nil {
		return nil, err
	}
	return &frontend.SeckillCouponKillRes{}, nil
}
