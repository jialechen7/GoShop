package captcha

import (
	"image/color"

	"github.com/afocus/captcha"
)

var cap *captcha.Captcha

func init() {
	cap = captcha.New()
	if err := cap.SetFont("comic.ttf"); err != nil {
		panic(err.Error())
	}

	cap.SetSize(80, 28)
	cap.SetDisturbance(captcha.NORMAL)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
}

func CreateCaptcha(len int) (*captcha.Image, string) {
	return cap.Create(len, captcha.ALL)
}
