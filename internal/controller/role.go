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

func (c *cRole) AddPermissions(ctx context.Context, req *backend.RoleAddPermissionsReq) (res *backend.RoleAddPermissionsRes, err error) {
	err = service.Role().AddPermissions(ctx, model.RoleAddPermissionsInput{
		RoleId:        req.RoleId,
		PermissionIds: req.PermissionIds,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (c *cRole) DeletePermissions(ctx context.Context, req *backend.RoleDeletePermissionsReq) (res *backend.RoleDeletePermissionsRes, err error) {
	err = service.Role().DeletePermissions(ctx, model.RoleDeletePermissionsInput{
		RoleId:        req.RoleId,
		PermissionIds: req.PermissionIds,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (c *cRole) Delete(ctx context.Context, req *backend.RoleDeleteReq) (res *backend.RoleDeleteRes, err error) {
	err = service.Role().Delete(ctx, req.Id)
	return
}

func (c *cRole) Update(ctx context.Context, req *backend.RoleUpdateReq) (res *backend.RoleUpdateRes, err error) {
	err = service.Role().Update(ctx, model.RoleUpdateInput{
		Id: req.Id,
		RoleCreateUpdateBase: model.RoleCreateUpdateBase{
			Name: req.Name,
			Desc: req.Desc,
		},
	})
	return
}

func (c *cRole) List(ctx context.Context, req *backend.RoleGetListCommonReq) (res *backend.RoleGetListCommonRes, err error) {
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
