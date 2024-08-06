package api

import "gin_gorm/api/captcha"

type AppGroup struct {
	CaptchaApi captcha.Captcha
}

var AppGroupApp = new(AppGroup)
