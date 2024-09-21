package api

import (
	"template/api/captcha"
	"template/api/chat"
)

type AppGroup struct {
	CaptchaApi captcha.Captcha
	ChatApi    chat.Chat
}

var AppGroupApp = new(AppGroup)
