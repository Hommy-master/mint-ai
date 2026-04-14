package config

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/silenceper/wechat/v2/cache"
	wechatConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"gopkg.in/yaml.v3"
)

var (
	cfg   = &Config{}
	weCfg = &wechatConfig.Config{} // 微信配置
	once  sync.Once
)

// 环境变量
var (
	// 生成资源的下载地址
	URL = os.Getenv("URL")

	// 微信扫码登录场景ID
	SceneID = os.Getenv("SceneID")

	// 微信公众号/服务号
	WeChatAPPID          = os.Getenv("WeChatAPPID")
	WeChatAPPSecret      = os.Getenv("WeChatAPPSecret")
	WeChatToken          = os.Getenv("WeChatToken")
	WeChatEncodingAESKey = os.Getenv("WeChatEncodingAESKey")
	WeChatUseStableAK    = os.Getenv("WeChatUseStableAK") == "true"

	// 微信支付
	WeChatMchID             = os.Getenv("WeChatMchID")
	WeChatMchAPIv3Key       = os.Getenv("WeChatMchAPIv3Key")
	WeChatMchCertSN         = os.Getenv("WeChatMchCertSN")
	WeChatPayNotifyURL      = URL + "/openapi/v1/user/pay/wechat-notify" // 微信支付 - 通知地址
	WeChatMchPrivateKeyPath = "/app/etc/wechat_pay_key.pem"

	// 数据库连接信息
	MySQL      = os.Getenv("MySQL")
	PostgreSQL = os.Getenv("PostgreSQL")

	// JWT签名密钥
	JWTSigningKey = os.Getenv("JWT_SIGNING_KEY")

	// JWT过期时间（天）
	JWTExpireDays = os.Getenv("JWT_EXPIRE_DAYS")
)

type Config struct {
	Server Server `yaml:"server"`
	Logger Logger `yaml:"logger"`
	FFmpeg FFmpeg `yaml:"ffmpeg"`
}

type Server struct {
	Addr              string `yaml:"addr"`
	OpenApiPath       string `yaml:"openApiPath"`
	SwaggerPath       string `yaml:"swaggerPath"`
	ClientMaxBodySize string `yaml:"clientMaxBodySize"`
	Logger            Logger `yaml:"logger"`
}

type Logger struct {
	Path                 string        `yaml:"path"`                 // 日志文件路径，可选值为：stdout, stderr, /path/to/logfile, off
	Level                string        `yaml:"level"`                // 日志级别，可选值为：debug, info, warning, error, fatal, panic, all, off
	Stdout               bool          `yaml:"stdout"`               // 是否输出到标准输出，可选值为：true, false
	RotateExpire         time.Duration `yaml:"rotateExpire"`         // 删除超过多长时间的切分文件
	RotateBackupExpire   time.Duration `yaml:"rotateBackupExpire"`   // 多长时间删除一次备份日志
	RotateBackupLimit    int           `yaml:"rotateBackupLimit"`    // 保留的备份日志文件数量
	RotateBackupCompress int           `yaml:"rotateBackupCompress"` // 压缩等级，默认为9，取值范围0-9，0表示不压缩，9表示最高压缩
}

type FFmpeg struct {
	Workspace    string `yaml:"workspace"`
	UploadOutput string `yaml:"uploadOutput"`
	Bots         string `yaml:"bots"`
}

type Fonts struct {
	Default string `yaml:"default"` // 默认字体文件路径
	Dir     string `yaml:"dir"`     // 字体文件目录
}

func check(c *Config) error {
	if c.Server.Addr == "" {
		return fmt.Errorf("server.addr is empty")
	}

	if c.Server.OpenApiPath == "" {
		return fmt.Errorf("server.openApiPath is empty")
	}

	if c.Server.SwaggerPath == "" {
		return fmt.Errorf("server.swaggerPath is empty")
	}

	if c.Server.Logger.Path == "" {
		return fmt.Errorf("server.logger.path is empty")
	}

	if c.Logger.Path == "" {
		return fmt.Errorf("logger.path is empty")
	}

	return nil
}

// 读取文件并解析为Config结构体
func Init(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return err
	}

	if err = check(cfg); err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	return cfg
}

func GetWeChatConfig() *wechatConfig.Config {
	once.Do(func() {
		weCfg = &wechatConfig.Config{
			AppID:          WeChatAPPID,
			AppSecret:      WeChatAPPSecret,
			Token:          WeChatToken,
			EncodingAESKey: WeChatEncodingAESKey,
			UseStableAK:    WeChatUseStableAK,
			Cache:          cache.NewMemory(),
		}
	})

	return weCfg
}
