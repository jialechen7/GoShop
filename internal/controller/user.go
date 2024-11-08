package controller

import (
	"context"
	"goshop/api/backend"
	"goshop/api/frontend"
	"goshop/internal/consts"

	"github.com/gogf/gf/util/gconv"

	"goshop/internal/model"
	"goshop/internal/service"
)

// User 用户管理
var User = cUser{}

type cUser struct{}

func (c *cUser) Create(ctx context.Context, req *frontend.UserRegisterReq) (res *frontend.UserRegisterRes, err error) {
	out, err := service.User().Create(ctx, model.UserCreateInput{
		UserCreateUpdateBase: model.UserCreateUpdateBase{
			Name:         req.Name,
			Password:     req.Password,
			Avatar:       req.Avatar,
			Sign:         req.Sign,
			Sex:          req.Sex,
			SecretAnswer: req.SecretAnswer,
		},
	})
	if err != nil {
		return nil, err
	}
	return &frontend.UserRegisterRes{UserId: out.Id}, nil
}

func (c *cUser) Info(ctx context.Context, req *frontend.UserGetInfoReq) (res *frontend.UserGetInfoRes, err error) {
	return &frontend.UserGetInfoRes{
		Id:     gconv.Int(ctx.Value(consts.CtxUserId)),
		Name:   gconv.String(ctx.Value(consts.CtxUserName)),
		Avatar: gconv.String(ctx.Value(consts.CtxUserAvatar)),
		Sign:   gconv.String(ctx.Value(consts.CtxUserSign)),
		Sex:    gconv.Int(ctx.Value(consts.CtxUserSex)),
		Status: gconv.Int(ctx.Value(consts.CtxUserStatus)),
	}, nil
}

// ResetPassword 重置密码（用户自己）
func (c *cUser) ResetPassword(ctx context.Context, req *frontend.UserResetPasswordReq) (res *frontend.UserResetPasswordRes, err error) {
	_, err = service.User().ResetPassword(ctx, model.UserUpdateInput{
		UserCreateUpdateBase: model.UserCreateUpdateBase{
			Password:     req.Password,
			SecretAnswer: req.SecretAnswer,
		},
		Id: gconv.Int(ctx.Value(consts.CtxUserId)),
	})
	return res, err
}

// Update 更新用户（管理员）
func (c *cUser) Update(ctx context.Context, req *backend.UserUpdateReq) (res *backend.UserUpdateRes, err error) {
	_, err = service.User().Update(ctx, model.UserUpdateInput{
		Id: req.Id,
		UserCreateUpdateBase: model.UserCreateUpdateBase{
			Name:         req.Name,
			Password:     req.Password,
			Avatar:       req.Avatar,
			Sex:          req.Sex,
			Sign:         req.Sign,
			SecretAnswer: req.SecretAnswer,
			Status:       req.Status,
		},
	})
	if err != nil {
		return nil, err
	}
	return res, err
}

func (c *cUser) List(ctx context.Context, req *backend.UserGetListCommonReq) (res *backend.UserGetListCommonRes, err error) {
	getListRes, err := service.User().GetList(ctx, model.UserGetListInput{
		Page: req.Page,
		Size: req.Size,
	})
	if err != nil {
		return nil, err
	}
	return &backend.UserGetListCommonRes{
		List:  getListRes.List,
		Page:  getListRes.Page,
		Size:  getListRes.Size,
		Total: getListRes.Total,
	}, nil
}

func (c *cUser) Delete(ctx context.Context, req *backend.UserDeleteReq) (res *backend.UserDeleteRes, err error) {
	err = service.User().Delete(ctx, req.Id)
	return
}
