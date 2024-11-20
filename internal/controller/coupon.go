package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/internal/model"
	"goshop/internal/service"
)

// Coupon 优惠券管理
var Coupon = cCoupon{}

type cCoupon struct{}

// List 查询优惠券列表
func (c *cCoupon) List(ctx context.Context, req *backend.CouponGetListCommonReq) (res *backend.CouponGetListCommonRes, err error) {
	getListRes, err := service.Coupon().GetList(ctx, model.CouponGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.CouponGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cCoupon) Add(ctx context.Context, req *backend.CouponAddReq) (res *backend.CouponAddRes, err error) {
	out, err := service.Coupon().Add(ctx, model.CouponAddInput{
		CouponCreateUpdateBase: model.CouponCreateUpdateBase{
			Name:       req.Name,
			Price:      req.Price,
			GoodsIds:   req.GoodsIds,
			CategoryId: req.CategoryId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.CouponAddRes{
		CouponId: out.CouponId,
	}, nil
}

func (c *cCoupon) Delete(ctx context.Context, req *backend.CouponDeleteReq) (res *backend.CouponDeleteRes, err error) {
	err = service.Coupon().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &backend.CouponDeleteRes{}, nil
}

func (c *cCoupon) Update(ctx context.Context, req *backend.CouponUpdateReq) (res *backend.CouponUpdateRes, err error) {
	err = service.Coupon().Update(ctx, model.CouponUpdateInput{
		Id: req.Id,
		CouponCreateUpdateBase: model.CouponCreateUpdateBase{
			Name:       req.Name,
			Price:      req.Price,
			GoodsIds:   req.GoodsIds,
			CategoryId: req.CategoryId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.CouponUpdateRes{}, nil
}
