package webserver

import (
	"context"
	"cozeos/internal/config"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"

	_ "cozeos/internal/pkg/validator"
)

func parseSizeString(sizeStr string) (int64, error) {
	re := regexp.MustCompile(`^(\d+)([KMGT]?B?)$`)
	matches := re.FindStringSubmatch(strings.ToUpper(sizeStr))
	if len(matches) != 3 {
		return 0, fmt.Errorf("invalid size format: %s", sizeStr)
	}

	value, err := strconv.ParseInt(matches[1], 10, 64)
	if err != nil {
		return 0, err
	}

	unit := matches[2]
	switch unit {
	case "B", "":
		return value, nil
	case "KB", "K":
		return value * 1024, nil
	case "MB", "M":
		return value * 1024 * 1024, nil
	case "GB", "G":
		return value * 1024 * 1024 * 1024, nil
	case "TB", "T":
		return value * 1024 * 1024 * 1024 * 1024, nil
	default:
		return 0, fmt.Errorf("unsupported size unit: %s", unit)
	}
}

func initLog(s *ghttp.Server) {
	// 服务日志配置
	s.SetLogPath(config.GetConfig().Server.Logger.Path)                // 日志目录
	s.SetLogStdout(config.GetConfig().Server.Logger.Stdout)            // HTTP访问日志输出到控制台（如：用Postman发送一个HTTP请求记录）
	s.Logger().SetStdoutPrint(config.GetConfig().Server.Logger.Stdout) // 服务启动日志输出到控制台（如：路由信息等）

	// 全局glog日志配置
	glog.SetPath(config.GetConfig().Logger.Path)
	glog.SetFile("{Y-m-d}.log")
	if config.GetConfig().Logger.Level == "PROD" {
		glog.SetLevel(glog.LEVEL_PROD)
	} else {
		glog.SetLevel(glog.LEVEL_DEV)
	}
	glog.SetStdoutPrint(config.GetConfig().Logger.Stdout)
	glog.SetFlags(glog.F_FILE_SHORT | glog.F_TIME_TIME)
	logCfg := glog.DefaultLogger().GetConfig()
	logCfg.RotateExpire = config.GetConfig().Logger.RotateExpire
	logCfg.RotateBackupLimit = config.GetConfig().Logger.RotateBackupLimit
	logCfg.RotateBackupExpire = config.GetConfig().Logger.RotateBackupExpire
	logCfg.RotateBackupCompress = config.GetConfig().Logger.RotateBackupCompress
	glog.SetConfig(logCfg)
}

func start(ctx context.Context, cfgPath string) error {
	// 初始化配置
	err := config.Init(cfgPath)
	if err != nil {
		glog.Errorf(ctx, "config.Init failed: %v", err)
		return err
	}
	fmt.Printf("config: %+v\n", config.GetConfig())
	fmt.Printf("URL: %s\n", config.URL)
	fmt.Printf("WeChatAPPID: %s\n", config.WeChatAPPID)
	fmt.Printf("WeChatAPPSecret: %s\n", config.WeChatAPPSecret)
	fmt.Printf("WeChatToken: %s\n", config.WeChatToken)
	fmt.Printf("WeChatEncodingAESKey: %s\n", config.WeChatEncodingAESKey)
	fmt.Printf("WeChatUseStableAK: %t\n", config.WeChatUseStableAK)
	fmt.Printf("clientMaxBodySize: %s\n", config.GetConfig().Server.ClientMaxBodySize)

	s := g.Server()
	// 设置服务地址
	s.SetAddr(config.GetConfig().Server.Addr)
	// 设置最大请求体大小
	if config.GetConfig().Server.ClientMaxBodySize != "" {
		// 解析大小字符串并转换为字节数
		sizeInBytes, err := parseSizeString(config.GetConfig().Server.ClientMaxBodySize)
		if err == nil {
			s.SetClientMaxBodySize(sizeInBytes)
			fmt.Printf("SetClientMaxBodySize: %s -> %d\n", config.GetConfig().Server.ClientMaxBodySize, sizeInBytes)
		}
	}
	// 初始化路由
	initRouter(ctx, s)
	// 初始化日志
	initLog(s)
	// 文档生成
	s.SetOpenApiPath(config.GetConfig().Server.OpenApiPath)
	s.SetSwaggerPath(config.GetConfig().Server.SwaggerPath)

	// 初始化定时任务
	initGCron(ctx)

	// 启动服务
	glog.Infof(ctx, "start server")

	s.Run()
	return nil
}
