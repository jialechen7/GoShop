package session

import (
	"context"
	"goshop/internal/consts"
	"goshop/internal/model/entity"
	"goshop/internal/service"
)

type sSession struct{}

func init() {
	service.RegisterSession(New())
}

func New() *sSession {
	return &sSession{}
}

// 设置管理员Session.
func (s *sSession) SetAdmin(ctx context.Context, admin *entity.AdminInfo) error {
	return service.BizCtx().Get(ctx).Session.Set(consts.SessionKeyAdmin, admin)
}

// 获取当前登录的管理员信息对象，如果管理员未登录返回nil。
func (s *sSession) GetAdmin(ctx context.Context) *entity.AdminInfo {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		v, _ := customCtx.Session.Get(consts.SessionKeyAdmin)
		if !v.IsNil() {
			var admin *entity.AdminInfo
			_ = v.Struct(&admin)
			return admin
		}
	}
	return &entity.AdminInfo{}
}

// 删除管理员Session。
func (s *sSession) RemoveAdmin(ctx context.Context) error {
	customCtx := service.BizCtx().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(consts.SessionKeyAdmin)
	}
	return nil
}
