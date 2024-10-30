// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goshop/internal/model/entity"
)

type (
	ISession interface {
		// SetAdmin 设置管理员Session
		SetAdmin(ctx context.Context, admin *entity.AdminInfo) error
		// GetAdmin 获取当前登录的管理员信息对象，如果管理员未登录返回nil
		GetAdmin(ctx context.Context) *entity.AdminInfo
		// RemoveAdmin 删除管理员Session
		RemoveAdmin(ctx context.Context) error
	}
)

var (
	localSession ISession
)

func Session() ISession {
	if localSession == nil {
		panic("implement not found for interface ISession, forgot register?")
	}
	return localSession
}

func RegisterSession(i ISession) {
	localSession = i
}
