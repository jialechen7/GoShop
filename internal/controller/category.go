package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/api/frontend"
	"goshop/internal/model"
	"goshop/internal/service"
)

// Category 分类管理
var Category = cCategory{}

type cCategory struct{}

// List 查询分类列表
func (c *cCategory) List(ctx context.Context, req *backend.CategoryGetListCommonReq) (res *backend.CategoryGetListCommonRes, err error) {
	getListRes, err := service.Category().GetList(ctx, model.CategoryGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.CategoryGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

// ListWithParentId 查询分类列表
func (c *cCategory) ListWithParentId(ctx context.Context, req *frontend.CategoryGetListCommonReq) (res *frontend.CategoryGetListCommonRes, err error) {
	getListRes, err := service.Category().GetListFrontend(ctx, model.CategoryGetListWithParentIdInput{
		Page:     req.Page,
		Size:     req.Size,
		ParentId: req.ParentId,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.CategoryGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

// ListAll 获取全部一级分类
func (c *cCategory) ListAll(ctx context.Context, req *backend.CategoryGetAllListCommonReq) (res *backend.CategoryGetAllListCommonRes, err error) {
	getAllRes, err := service.Category().GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return &backend.CategoryGetAllListCommonRes{
		List:  getAllRes.List,
		Total: getAllRes.Total,
	}, nil
}

func (c *cCategory) Add(ctx context.Context, req *backend.CategoryAddReq) (res *backend.CategoryAddRes, err error) {
	out, err := service.Category().Add(ctx, model.CategoryAddInput{
		CategoryCreateUpdateBase: model.CategoryCreateUpdateBase{
			Name:     req.Name,
			PicUrl:   req.PicUrl,
			Level:    req.Level,
			Sort:     req.Sort,
			ParentId: req.ParentId,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.CategoryAddRes{
		CategoryId: out.CategoryId,
	}, nil
}

func (c *cCategory) Delete(ctx context.Context, req *backend.CategoryDeleteReq) (res *backend.CategoryDeleteRes, err error) {
	err = service.Category().Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &backend.CategoryDeleteRes{}, nil
}
