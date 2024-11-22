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
	captachaImage, captachaCode := captcha.CreateCaptcha(consts.CaptachaLength)
	g.Dump(captachaCode)
	_, err = g.Redis().Do(ctx, "SET", consts.CaptachaPrefix+captchaId, captachaCode, "EX", consts.CaptachaExpire)
	if err != nil {
		return "", nil, err
	}
	emptyBuffer := bytes.NewBuffer(nil)
	err = png.Encode(emptyBuffer, captachaImage)
	if err != nil {
		return "", nil, err
	}
	return captchaId, emptyBuffer.Bytes(), nil
}
