package backend

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginDoReq struct {
	g.Meta   `path:"/backend/login" method:"post" summary:"执行登录请求" tags:"登录"`
	Name     string `json:"name" v:"required#请输入用户名"   dc:"用户名"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
}
type LoginDoRes struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expire_at"`
}

type RefreshTokenReq struct {
	g.Meta `path:"/backend/refresh_token" method:"post"`
}

type RefreshTokenRes struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expire_at"`
}

type LogoutReq struct {
	g.Meta `path:"/backend/logout" method:"post"`
}

type LogoutRes struct {
}
