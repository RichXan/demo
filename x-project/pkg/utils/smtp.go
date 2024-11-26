package utils

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

const (
	HTML  = "html"
	PLAIN = "plain"
)

// EmailClient 连接smtp需要的参数
type EmailClient struct {
	Host     string
	Port     int
	UserName string // 认证则需要传
	PassWord string // 认证则需要传
}

// EmailParams 发送邮件需要的参数
type EmailParams struct {
	From     string   // 发件邮箱
	To       []string // 收件人
	Cc       []string // 抄送
	Bcc      []string // 密件抄送
	Subject  string   // 标题
	Body     string   // 内容
	BodyType string   // 内容类型
	Attach   string   // 附件
}

// SendEmail 发送邮件
func SendEmail(client EmailClient, params EmailParams) error {
	m := gomail.NewMessage()
	m.SetHeader("From", params.From)       // 发件人
	m.SetHeader("To", params.To...)        // 收件人
	m.SetHeader("Cc", params.Cc...)        // 抄送人
	m.SetHeader("Bcc", params.Bcc...)      // 抄送人
	m.SetHeader("Subject", params.Subject) // 邮件主题
	switch params.BodyType {
	case HTML:
		m.SetBody("text/html", params.Body)
	case PLAIN:
		m.SetBody("text/plain", params.Body)
	default:
		break
	}
	if params.Attach != "" {
		m.Attach(params.Attach)
	}

	d := gomail.NewDialer(
		client.Host,
		client.Port,
		client.UserName,
		client.PassWord,
	)
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
