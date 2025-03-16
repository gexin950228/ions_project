package utils

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type Captcha struct {
	Id   string
	BS64 string
	Code int
}

var store = base64Captcha.DefaultMemStore

func GetCaptcha() (id, base64 string, err error) {
	rgbaColor := color.RGBA{0, 0, 0, 0}
	fonts := []string{"wqy-microhei.ttc"}
	driver := base64Captcha.NewDriverMath(50, 140, 0, 0, &rgbaColor, fonts)
	captcha := base64Captcha.NewCaptcha(driver, store)

	id, base64, err = captcha.Generate()
	return id, base64, err
}

func VerifyCaptcha(id, code string) bool {
	return store.Verify(id, code, true)
}
