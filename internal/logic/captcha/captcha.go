package captcha

import (
	"bytes"
	"context"
	"goshop/internal/consts"
	"goshop/internal/service"
	"goshop/utility/captcha"
	"image/png"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type sCaptcha struct{}

func init() {
	service.RegisterCaptcha(New())
}

func New() *sCaptcha {
	return &sCaptcha{}
}

func (s *sCaptcha) GetCaptcha(ctx context.Context) (captchaId string, image []byte, err error) {
	captchaId = uuid.NewString()
	captchaImage, captchaCode := captcha.CreateCaptcha(consts.CaptchaLength)
	g.Dump(captchaId, captchaCode)
	_, err = g.Redis().Do(ctx, "SET", consts.CaptchaPrefix+captchaId, captchaCode, "EX", consts.CaptchaExpire)
	if err != nil {
		return "", nil, err
	}
	emptyBuffer := bytes.NewBuffer(nil)
	err = png.Encode(emptyBuffer, captchaImage)
	if err != nil {
		return "", nil, err
	}
	return captchaId, emptyBuffer.Bytes(), nil
}
