package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/api/frontend"
	"goshop/internal/model"
	"goshop/internal/service"
)

// Coupon 优惠券管理
var Coupon = cCoupon{}

type cCoupon struct{}

// ListBackend 查询优惠券列表
func (c *cCoupon) ListBackend(ctx context.Context, req *backend.CouponGetListCommonReq) (res *backend.CouponGetListCommonRes, err error) {
	getListRes, err := service.Coupon().GetListBackend(ctx, model.CouponGetListInput{
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
			Condition:  req.Condition,
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

// ListFrontend 查询优惠券列表
func (c *cCoupon) ListFrontend(ctx context.Context, req *frontend.CouponGetListCommonReq) (res *frontend.CouponGetListCommonRes, err error) {
	getListRes, err := service.Coupon().GetListFrontend(ctx, model.CouponGetListWithGoodsIdInput{
		GoodsId: req.GoodsId,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.CouponGetListCommonRes{
		List: getListRes.List,
	}, nil
}

func (c *cCoupon) ListAvailableFrontend(ctx context.Context, req *frontend.CouponGetListAvailableReq) (res *frontend.CouponGetListAvailableRes, err error) {
	getListRes, err := service.Coupon().GetListAvailable(ctx, model.CouponGetListAvailableInput{
		OrderGoodsInfos: req.OrderGoodsInfos,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.CouponGetListAvailableRes{
		AvailableList:   getListRes.AvailableList,
		UnavailableList: getListRes.UnavailableList,
	}, nil
}
