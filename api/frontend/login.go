package frontend

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// LoginRes 用于gtoken登录请求的参数
type LoginRes struct {
	Type     string `json:"type"`
	Token    string `json:"token"`
	ExpireIn int    `json:"expire_in"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	Sign     string `json:"sign"`
	Sex      int    `json:"sex"`
	Status   int    `json:"status"`
}

type RefreshTokenReq struct {
	g.Meta `path:"/refresh_token" method:"post"`
}

type RefreshTokenRes struct {
	Token    string    `json:"token"`
	ExpireIn time.Time `json:"expire_in"`
}

type LogoutReq struct {
	g.Meta `path:"/user/logout" method:"post"`
}

type LogoutRes struct {
}
