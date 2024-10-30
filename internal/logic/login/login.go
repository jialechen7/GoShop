package login

import (
	"context"
	"goshop/internal/dao"
	"goshop/internal/model"
	"goshop/internal/model/entity"
	"goshop/internal/service"
	"goshop/utility"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}

// Login 执行登录
func (s *sLogin) Login(ctx context.Context, in model.AdminLoginInput) error {
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return gerror.New("用户名不存在")
	}
	if utility.EncryptPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return gerror.New("密码不正确")
	}
	if err := service.Session().SetAdmin(ctx, &adminInfo); err != nil {
		return err
	}
	service.BizCtx().SetAdmin(ctx, &model.ContextAdmin{
		Id:      adminInfo.Id,
		Name:    adminInfo.Name,
		RoleIds: adminInfo.RoleIds,
		IsAdmin: adminInfo.IsAdmin,
	})
	return nil
}
