// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goshop/internal/model"
)

type (
	IOauth interface {
		GithubLogin(ctx context.Context) (string, error)
		// GithubReceiveCode 接收github回调并自动登录（未注册自动注册）
		GithubReceiveCode(ctx context.Context, in model.GithubReceiveCodeInput) (out model.GithubReceiveCodeOutput, err error)
	}
)

var (
	localOauth IOauth
)

func Oauth() IOauth {
	if localOauth == nil {
		panic("implement not found for interface IOauth, forgot register?")
	}
	return localOauth
}

func RegisterOauth(i IOauth) {
	localOauth = i
}
