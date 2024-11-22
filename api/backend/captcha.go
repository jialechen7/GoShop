package backend

import "github.com/gogf/gf/v2/frame/g"

type GetCaptchaReq struct {
	g.Meta `path:"/captcha/get" method:"get" tags:"获取验证码后台" summary:"获取验证码接口"`
}

type GetCaptchaRes struct {
	CaptchaId string `json:"captcha_id"`
	Image     string `json:"image"`
}
