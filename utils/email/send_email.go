package email

import (
	"context"
	"gopkg.in/gomail.v2"
	"template/global"
	"template/utils/random"
	"time"
)

type Subject string

const (
	Code Subject = "验证码"
)

type Api struct {
	Subject Subject
}

func (a Api) Send(name string) error {
	return send(name, string(a.Subject))
}

func NewCode() Api {
	return Api{
		Subject: Code,
	}
}

// send 邮件发送  发给谁，主题，正文
func send(name, subject string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	e := global.Config.Email
	code := random.Code()
	body := "你的验证码是" + code + "有效期1分钟"
	global.Redis.Set(ctx, name, code, time.Minute)
	return sendMail(
		e.User,
		e.Password,
		e.Host,
		e.Port,
		name,
		e.DefaultFromEmail,
		subject,
		body,
	)
}

func sendMail(userName, authCode, host string, port int, mailTo, sendName string, subject, body string) error {
	//创建一个邮件消息对象
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, sendName)) // 发件人邮箱，发件人名字
	m.SetHeader("To", mailTo)                                // 发送给谁
	m.SetHeader("Subject", subject)                          // 主题
	m.SetBody("text/html", body)
	/*
		host: 邮件服务器的主机地址
		port: 邮件服务器的端口号
		userName: 发件人的邮箱账号
		authCode: 发件人的邮箱授权码
	*/
	d := gomail.NewDialer(host, port, userName, authCode)
	//发送
	err := d.DialAndSend(m)
	return err
}
