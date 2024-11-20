package controller

import (
	"context"
	"goshop/api/frontend"
	"goshop/internal/model"
	"goshop/internal/service"
)

// UserCoupon 优惠券管理
var UserCoupon = cUserCoupon{}

type cUserCoupon struct{}

// List 查询优惠券列表
func (c *cUserCoupon) List(ctx context.Context, req *frontend.UserCouponGetListCommonReq) (res *frontend.UserCouponGetListCommonRes, err error) {
	getListRes, err := service.UserCoupon().GetList(ctx, model.UserCouponGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.UserCouponGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cUserCoupon) Add(ctx context.Context, req *frontend.UserCouponAddReq) (res *frontend.UserCouponAddRes, err error) {
	out, err := service.UserCoupon().Add(ctx, model.UserCouponAddInput{
		UserCouponCreateUpdateBase: model.UserCouponCreateUpdateBase{
			CouponId: req.CouponId,
			Status:   req.Status,
		},
	})
	if err != nil {
		return nil, err
	}
	return &frontend.UserCouponAddRes{
		UserCouponId: out.UserCouponId,
	}, nil
}

func (c *cUserCoupon) Delete(ctx context.Context, req *frontend.UserCouponDeleteReq) (res *frontend.UserCouponDeleteRes, err error) {
	err = service.UserCoupon().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &frontend.UserCouponDeleteRes{}, nil
}

func (c *cUserCoupon) Update(ctx context.Context, req *frontend.UserCouponUpdateReq) (res *frontend.UserCouponUpdateRes, err error) {
	err = service.UserCoupon().Update(ctx, model.UserCouponUpdateInput{
		Id: req.Id,
		UserCouponCreateUpdateBase: model.UserCouponCreateUpdateBase{
			CouponId: req.CouponId,
			Status:   req.Status,
		},
	})
	if err != nil {
		return nil, err
	}
	return &frontend.UserCouponUpdateRes{}, nil
}
