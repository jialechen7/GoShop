package frontend

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserRegisterReq struct {
	g.Meta       `path:"/user/register" tags:"User" method:"post" summary:"用户注册接口"`
	Name         string `json:"name" v:"required#用户名不能为空" dc:"用户名"`
	Password     string `json:"password" v:"required#密码不能为空" dc:"密码"`
	Avatar       string `json:"avatar" dc:"头像"`
	Sign         string `json:"sign" dc:"个性签名"`
	Sex          int    `json:"sex" dc:"性别"`
	SecretAnswer string `json:"secret_answer" dc:"密保问题答案"`
}

type UserRegisterRes struct {
	UserId int `json:"userId"`
}

type UserDeleteReq struct {
	g.Meta `path:"/user/delete" method:"delete" tags:"User" summary:"删除用户接口"`
	Id     int `v:"min:1#请选择需要删除的用户" dc:"用户id"`
}
type UserDeleteRes struct{}

type UserGetInfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"用户" summary:"获取用户信息接口"`
}

// UserGetInfoRes 用于gtoken
type UserGetInfoRes struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Sign   string `json:"sign"`
	Sex    int    `json:"sex"`
	Status int    `json:"status"`
}

type UserResetPasswordReq struct {
	g.Meta       `path:"/sso/password/update" method:"post" tags:"用户" summary:"重置密码接口"`
	Password     string `json:"password" v:"required#密码不能为空" dc:"密码"`
	SecretAnswer string `json:"secret_answer" v:"required#密保问题答案不能为空" dc:"密保问题答案"`
}

type UserResetPasswordRes struct{}
