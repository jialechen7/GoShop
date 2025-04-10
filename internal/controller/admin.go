package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/internal/consts"

	"github.com/gogf/gf/util/gconv"

	"goshop/internal/model"
	"goshop/internal/service"
)

// Admin 管理员管理
var Admin = cAdmin{}

type cAdmin struct{}

func (c *cAdmin) List(ctx context.Context, req *backend.AdminGetListCommonReq) (res *backend.AdminGetListCommonRes, err error) {
	getListRes, err := service.Admin().GetList(ctx, model.AdminGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.AdminGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cAdmin) Create(ctx context.Context, req *backend.AdminReq) (res *backend.AdminRes, err error) {
	out, err := service.Admin().Create(ctx, model.AdminCreateInput{
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
	})
	if err != nil {
		return nil, err
	}
	return &backend.AdminRes{AdminId: out.AdminId}, nil
}

func (c *cAdmin) Delete(ctx context.Context, req *backend.AdminDeleteReq) (res *backend.AdminDeleteRes, err error) {
	err = service.Admin().Delete(ctx, req.Id)
	return
}

func (c *cAdmin) Update(ctx context.Context, req *backend.AdminUpdateReq) (res *backend.AdminUpdateRes, err error) {
	err = service.Admin().Update(ctx, model.AdminUpdateInput{
		Id: req.Id,
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Name:     req.Name,
			Password: req.Password,
			RoleIds:  req.RoleIds,
			IsAdmin:  req.IsAdmin,
		},
	})
	return
}

func (c *cAdmin) UpdatePassword(ctx context.Context, req *backend.AdminUpdatePasswordReq) (res *backend.AdminUpdatePasswordRes, err error) {
	err = service.Admin().Update(ctx, model.AdminUpdateInput{
		AdminCreateUpdateBase: model.AdminCreateUpdateBase{
			Password: req.Password,
		},
	})
	if err != nil {
		return nil, err
	}
	return
}

// Info should be authenticated to view.
// It is the get admin data handler
func (c *cAdmin) Info(ctx context.Context, req *backend.AdminGetInfoReq) (res *backend.AdminGetInfoRes, err error) {
	return &backend.AdminGetInfoRes{
		Id:      gconv.Int(ctx.Value(consts.CtxAdminId)),
		Name:    gconv.String(ctx.Value(consts.CtxAdminName)),
		IsAdmin: gconv.Int(ctx.Value(consts.CtxAdminIsAdmin)),
		RoleIds: gconv.String(ctx.Value(consts.CtxAdminRoleIds)),
	}, nil
}
