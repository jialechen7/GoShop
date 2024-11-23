package backend

import "github.com/gogf/gf/v2/frame/g"

type GithubLoginReq struct {
	g.Meta `path:"/oauth/github" method:"get" tags:"Github登录后台" summary:"点击跳转到Github登录页面接口"`
}

type OauthLoginRes struct {
	Url string `json:"url"`
}

type GithubReceiveCodeReq struct {
	g.Meta `path:"/oauth/github/receive_code" method:"get" tags:"Github登录后台" summary:"接收Github返回的code回调接口"`
	Code   string `json:"code" query:"code" v:"required#code不能为空"`
	State  string `json:"state" query:"state" v:"required#state不能为空"`
}

type GithubReceiveCodeRes struct{}
