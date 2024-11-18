package user

import (
	"context"
	"goshop/internal/consts"
	"goshop/internal/model/entity"
	"goshop/utility"

	"github.com/gogf/gf/v2/database/gdb"
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

// ResetPassword 重置密码（用户自己）
func (s *sUser) ResetPassword(ctx context.Context, in model.UserUpdateInput) (out model.UserUpdateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	in.Id = gconv.Int(ctx.Value(consts.CtxUserId))
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

// Update 更新用户（用户）
func (s *sUser) Update(ctx context.Context, in model.UserUpdateInput) (out model.UserUpdateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}

	var userInfo entity.UserInfo
	if err = dao.UserInfo.Ctx(ctx).Where(dao.UserInfo.Columns().Id, in.Id).Scan(&userInfo); err != nil {
		return out, err
	}
	in.UserSalt = grand.S(10)
	in.Password = utility.EncryptPassword(in.Password, in.UserSalt)
	_, err = dao.UserInfo.Ctx(ctx).Data(in).OmitEmpty().Where(dao.UserInfo.Columns().Id, in.Id).Update()
	if err != nil {
		return out, err
	}
	return out, nil
}

// GetList 查询用户列表
func (s *sUser) GetList(ctx context.Context, in model.UserGetListInput) (out *model.UserGetListOutput, err error) {
	var (
		m = dao.UserInfo.Ctx(ctx)
	)
	out = &model.UserGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	listModel := m.Page(in.Page, in.Size)

	// 执行查询
	var list []*entity.UserInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}

	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

// Delete 删除用户
func (s *sUser) Delete(ctx context.Context, id int) error {
	return dao.UserInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除用户
		_, err := dao.UserInfo.Ctx(ctx).Where(dao.UserInfo.Columns().Id, id).Delete()
		return err
	})
}
