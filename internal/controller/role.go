package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/internal/model"
	"goshop/internal/service"
)

var Role = cRole{}

type cRole struct{}

func (c *cRole) Create(ctx context.Context, req *backend.RoleCreateReq) (res *backend.RoleCreateRes, err error) {
	out, err := service.Role().Create(ctx, model.RoleCreateInput{
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleCreateRes{RoleId: out.RoleId}, nil
}

func (c *cRole) AddPermission(ctx context.Context, req *backend.RoleAddPermissionReq) (res *backend.RoleAddPermissionRes, err error) {
	out, err := service.Role().AddPermission(ctx, model.RoleAddPermissionInput{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleAddPermissionRes{RolePermissionId: out.RolePermissionId}, nil
}

func (a *cRole) DeletePermission(ctx context.Context, req *backend.RoleDeletePermissionReq) (res *backend.RoleDeletePermissionRes, err error) {
	err = service.Role().DeletePermission(ctx, model.RoleDeletePermissionInput{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	})
	return
}

func (a *cRole) Delete(ctx context.Context, req *backend.RoleDeleteReq) (res *backend.RoleDeleteRes, err error) {
	err = service.Role().Delete(ctx, req.Id)
	return
}

func (a *cRole) Update(ctx context.Context, req *backend.RoleUpdateReq) (res *backend.RoleUpdateRes, err error) {
	err = service.Role().Update(ctx, model.RoleUpdateInput{
		Id: req.Id,
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	return
}

func (a *cRole) List(ctx context.Context, req *backend.RoleGetListCommonReq) (res *backend.RoleGetListCommonRes, err error) {
	getListRes, err := service.Role().GetList(ctx, model.RoleGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.RoleGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}
