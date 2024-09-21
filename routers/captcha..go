package routers

import "template/api"

func (router RouterGroup) CaptchaRouter() {
	captchaApi := api.AppGroupApp.CaptchaApi
	router.GET("captcha", captchaApi.Generate)
}
