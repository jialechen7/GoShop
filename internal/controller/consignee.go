package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/api/frontend"
	"goshop/internal/consts"
	"goshop/internal/model"
	"goshop/internal/service"

	"github.com/gogf/gf/util/gconv"
)

// Consignee 收货人管理
var Consignee = cConsignee{}

type cConsignee struct{}

// ListBackend 查询收货人列表
func (c *cConsignee) ListBackend(ctx context.Context, req *backend.ConsigneeGetListCommonReq) (res *backend.ConsigneeGetListCommonRes, err error) {
	getListRes, err := service.Consignee().GetListBackend(ctx, model.ConsigneeGetListBackendInput{
		Page:  req.Page,
		Size:  req.Size,
		Name:  req.Name,
		Phone: req.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &backend.ConsigneeGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cConsignee) UpdateBackend(ctx context.Context, req *backend.ConsigneeUpdateReq) (res *backend.ConsigneeUpdateRes, err error) {
	err = service.Consignee().UpdateBackend(ctx, model.ConsigneeUpdateInput{
		Id: req.Id,
		ConsigneeCreateUpdateBase: model.ConsigneeCreateUpdateBase{
			UserId:    gconv.Int(ctx.Value(consts.CtxUserId)),
			IsDefault: req.IsDefault,
			Name:      req.Name,
			Phone:     req.Phone,
			Province:  req.Province,
			City:      req.City,
			Town:      req.Town,
			Street:    req.Street,
			Detail:    req.Detail,
		},
	})
	return
}

// AddBackend 添加收货人
func (c *cConsignee) AddBackend(ctx context.Context, req *backend.ConsigneeAddReq) (res *backend.ConsigneeAddRes, err error) {
	out, err := service.Consignee().AddBackend(ctx, model.ConsigneeAddInput{
		ConsigneeCreateUpdateBase: model.ConsigneeCreateUpdateBase{
			UserId:    gconv.Int(ctx.Value(consts.CtxUserId)),
			IsDefault: req.IsDefault,
			Name:      req.Name,
			Phone:     req.Phone,
			Province:  req.Province,
			City:      req.City,
			Town:      req.Town,
			Street:    req.Street,
			Detail:    req.Detail,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.ConsigneeAddRes{
		ConsigneeId: out.ConsigneeId,
	}, nil
}

// ListFrontend 查询收货人列表（仅用户自己）
func (c *cConsignee) ListFrontend(ctx context.Context, req *frontend.ConsigneeGetListCommonReq) (res *frontend.ConsigneeGetListCommonRes, err error) {
	getListRes, err := service.Consignee().GetListFrontend(ctx, model.ConsigneeGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &frontend.ConsigneeGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cConsignee) DeleteBackend(ctx context.Context, req *backend.ConsigneeDeleteReq) (res *backend.ConsigneeDeleteRes, err error) {
	err = service.Consignee().DeleteBackend(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &backend.ConsigneeDeleteRes{}, nil
}

func (c *cConsignee) AddFrontend(ctx context.Context, req *frontend.ConsigneeAddReq) (res *frontend.ConsigneeAddRes, err error) {
	out, err := service.Consignee().AddFrontend(ctx, model.ConsigneeAddInput{
		ConsigneeCreateUpdateBase: model.ConsigneeCreateUpdateBase{
			UserId:    gconv.Int(ctx.Value(consts.CtxUserId)),
			IsDefault: req.IsDefault,
			Name:      req.Name,
			Phone:     req.Phone,
			Province:  req.Province,
			City:      req.City,
			Town:      req.Town,
			Street:    req.Street,
			Detail:    req.Detail,
		},
	})
	if err != nil {
		return nil, err
	}
	return &frontend.ConsigneeAddRes{
		ConsigneeId: out.ConsigneeId,
	}, nil
}

func (c *cConsignee) DeleteFrontend(ctx context.Context, req *frontend.ConsigneeDeleteReq) (res *frontend.ConsigneeDeleteRes, err error) {
	err = service.Consignee().DeleteFrontend(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &frontend.ConsigneeDeleteRes{}, nil
}

func (c *cConsignee) UpdateFrontend(ctx context.Context, req *frontend.ConsigneeUpdateReq) (res *frontend.ConsigneeUpdateRes, err error) {
	err = service.Consignee().UpdateFrontend(ctx, model.ConsigneeUpdateInput{
		Id: req.Id,
		ConsigneeCreateUpdateBase: model.ConsigneeCreateUpdateBase{
			UserId:    gconv.Int(ctx.Value(consts.CtxUserId)),
			IsDefault: req.IsDefault,
			Name:      req.Name,
			Phone:     req.Phone,
			Province:  req.Province,
			City:      req.City,
			Town:      req.Town,
			Street:    req.Street,
			Detail:    req.Detail,
		},
	})
	return
}
