package captcha

import (
	"gin_gorm/global"
	"gin_gorm/models/res"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

type CaptchaResponse struct {
	CaptchaID string `json:"captchaId"`
	PicPath   string `json:"picPath"`
}

func (Captcha *Captcha) CreateCaptcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(
		global.Config.Captcha.ImgHeight,
		global.Config.Captcha.ImgWidth,
		global.Config.Captcha.KeyLong,
		0.7,
		70,
	)
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		global.Log.Error("fail to generate the captcha", zap.Error(err))
		return
	}
	res.OkWithData(CaptchaResponse{
		CaptchaID: id,
		PicPath:   b64s,
	}, c)
}
