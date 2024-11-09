package backend

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// LoginDoReq 用于JWT登录请求的参数
type LoginDoReq struct {
	g.Meta   `path:"/login" method:"post" summary:"执行登录请求" tags:"登录"`
	Name     string `json:"name" v:"required#请输入用户名"   dc:"用户名"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
}

// LoginDoRes 用于JWT登录成功后返回的数据
type LoginDoRes struct {
	Token    string    `json:"token"`
	ExpireIn time.Time `json:"expire_in"`
}

// LoginRes 用于gtoken登录请求的参数
type LoginRes struct {
	Type        string      `json:"type"`
	Token       string      `json:"token"`
	ExpireIn    int         `json:"expire_in"`
	IsAdmin     int         `json:"is_admin"`
	RoleIds     string      `json:"role_ids"`
	Permissions interface{} `json:"permissions"`
}

type RefreshTokenReq struct {
	g.Meta `path:"/refresh_token" method:"post"`
}

type RefreshTokenRes struct {
	Token    string    `json:"token"`
	ExpireIn time.Time `json:"expire_in"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" method:"post"`
}

type LogoutRes struct {
}
