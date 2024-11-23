package controller

import (
	"context"
	"encoding/base64"
	"goshop/api/backend"
	"goshop/internal/service"
)

var Captcha = cCaptcha{}

type cCaptcha struct{}

func (c *cCaptcha) Get(ctx context.Context, req *backend.GetCaptchaReq) (res *backend.GetCaptchaRes, err error) {
	captchaId, image, err := service.Captcha().GetCaptcha(ctx)
	if err != nil {
		return nil, err
	}
	return &backend.GetCaptchaRes{
		CaptchaId: captchaId,
		Image:     "data:image/png;base64," + base64.StdEncoding.EncodeToString(image),
	}, nil
}
