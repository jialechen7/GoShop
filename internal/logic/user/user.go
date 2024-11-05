package user

import (
	"context"
	"goshop/internal/consts"
	"goshop/internal/model/entity"
	"goshop/utility"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/v2/util/grand"

	"goshop/internal/service"

	"github.com/gogf/gf/v2/encoding/ghtml"

	"goshop/internal/dao"
	"goshop/internal/model"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// Create 创建用户
func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (out model.UserCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}

	// 生成用户盐并加密密码
	in.UserSalt = grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, in.UserSalt)

	lastInsertID, err := dao.UserInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.UserCreateOutput{Id: int(lastInsertID)}, err
}

// ResetPassword 重置密码
func (s *sUser) ResetPassword(ctx context.Context, in model.UserUpdateInput) (out model.UserUpdateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}

	userStatus := gconv.Int(ctx.Value(consts.CtxUserStatus))
	if userStatus == consts.UserStatusBlacked {
		return out, gerror.New(consts.ErrUserStatus)
	}

	if len(in.Password) < consts.MinPasswordLength {
		return out, gerror.Newf("密码长度不能小于%d", consts.MinPasswordLength)
	}

	userId := gconv.Int(ctx.Value(consts.CtxUserId))
	var userInfo entity.UserInfo
	if err = dao.UserInfo.Ctx(ctx).Where(dao.UserInfo.Columns().Id, userId).Scan(&userInfo); err != nil {
		return out, err
	}
	if userInfo.SecretAnswer != in.SecretAnswer {
		return out, gerror.New(consts.ErrSecretAnswer)
	}

	in.UserSalt = grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, in.UserSalt)
	_, err = dao.UserInfo.Ctx(ctx).Data(in).OmitEmpty().Where(dao.UserInfo.Columns().Id, userId).Update()
	if err != nil {
		return out, err
	}
	return out, nil
}
