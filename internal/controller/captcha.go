package controller

import (
	"context"
	"encoding/base64"
	"goshop/api/backend"
	"goshop/internal/service"
)

var Captacha = cCaptacha{}

type cCaptacha struct{}

func (c *cCaptacha) Get(ctx context.Context, req *backend.GetCaptchaReq) (res *backend.GetCaptchaRes, err error) {
	captachaId, image, err := service.Captcha().GetCaptcha(ctx)
	if err != nil {
		return nil, err
	}
	return &backend.GetCaptchaRes{
		CaptchaId: captachaId,
		Image:     "data:image/png;base64," + base64.StdEncoding.EncodeToString(image),
	}, nil
}
