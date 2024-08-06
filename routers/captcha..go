package routers

import "gin_gorm/api"

func (router RouterGroup) CaptchaRouter() {
	captchaApi := api.AppGroupApp.CaptchaApi
	router.GET("captcha", captchaApi.CreateCaptcha)
}
