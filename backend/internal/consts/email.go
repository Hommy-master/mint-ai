package consts

import (
	"fmt"
	"os"
	"strconv"
)

// 通用邮件服务器定义，在服务中，所有需要发送邮件的地方，都使用环境变量配置
var (
	SmtpHost = mustGetEnv("SMTP_HOST")
	SmtpPort = mustGetEnvInt("SMTP_PORT")
	SmtpUser = mustGetEnv("SMTP_USER")
	SmtpPass = mustGetEnv("SMTP_PASS")
)

func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("missing required environment variable: %s", key))
	}
	return val
}

func mustGetEnvInt(key string) int {
	val := mustGetEnv(key)
	port, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Sprintf("invalid integer environment variable %s: %v", key, err))
	}
	return port
}
