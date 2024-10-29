package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/internal/model"
	"goshop/internal/service"
)

// 登录管理
var Login = cLogin{}

type cLogin struct{}

func (a *cLogin) Login(ctx context.Context, req *backend.LoginDoReq) (res *backend.LoginDoRes, err error) {
	res = &backend.LoginDoRes{}
	err = service.Login().Login(ctx, model.AdminLoginInput{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	loginAdmin := service.Session().GetAdmin(ctx)
	res.Info = loginAdmin
	return
}
