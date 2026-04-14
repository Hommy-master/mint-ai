package types

type SMTPServer struct {
	Host string `json:"host" v:"required|length:5,128"` // SMTP 服务器地址
	Port int    `json:"port" v:"between:1,65535"`       // SMTP 服务器端口
	User string `json:"user" v:"required|email"`        // SMTP 用户名（发件人邮箱）
	Pass string `json:"pass" v:"required|length:0,256"` // SMTP 密码（发件人邮箱的授权码）
}

type Email struct {
	To      []string `json:"to" v:"foreach|email"`       // 收件人邮箱地址列表
	Subject string   `json:"subject" v:"length:0,128"`   // 邮件主题
	Body    string   `json:"body" v:"length:0,10240000"` // 邮件正文内容
}
