package util

import (
	"gopkg.in/gomail.v2"
)

type BodyType int

const (
	PLAIN = BodyType(0)
	HTML  = BodyType(1)
)

// SendMail 发送邮件
// 参数：
// - smtpHost: SMTP 服务器地址
// - smtpPort: SMTP 服务器端口
// - smtpUser: SMTP 用户名（发件人邮箱）
// - smtpPass: SMTP 密码（发件人邮箱的授权码）
// - from: 发件人邮箱地址
// - to: 收件人邮箱地址列表
// - subject: 邮件主题
// - body: 邮件正文内容
// - bodyType: 邮件正文类型，"text/plain" 表示普通文本，"text/html" 表示 HTML 格式
// 返回值：
// - error: 发送邮件过程中遇到的错误，无错误则返回 nil
func SendMail(smtpHost string, smtpPort int, smtpUser string, smtpPass string,
	from string, to []string, subject string, body string, bodyType BodyType) error {
	// 创建邮件消息
	msg := gomail.NewMessage()

	// 设置发件人、收件人和主题
	msg.SetHeader("From", from)
	msg.SetHeader("To", to...)
	msg.SetHeader("Subject", subject)

	// 设置邮件正文
	switch bodyType {
	case PLAIN:
		msg.SetBody("text/plain", body)
	case HTML:
		msg.SetBody("text/html", body)
	default:
		msg.SetBody("text/plain", body)
	}

	// 创建拨号器
	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)

	// 发送邮件
	return dialer.DialAndSend(msg)
}
