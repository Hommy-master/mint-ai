package middleware

import (
	"cozeos/internal/config"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

// 创建工作目录
func (s *sMiddleware) mkdir(r *ghttp.Request) error {
	t := time.Now()
	dateStr := t.Format("2006-01-02") + "-" + fmt.Sprintf("%02d", t.Hour())
	workspace := config.GetConfig().FFmpeg.Workspace + "/" + dateStr
	uploadOutput := config.GetConfig().FFmpeg.UploadOutput + "/" + dateStr
	bots := config.GetConfig().FFmpeg.Bots // 智能体上传目录

	dirs := []string{
		workspace,
		uploadOutput,
	}

	// 检查目录是否存在，如果不存在则创建
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				glog.Errorf(r.Context(), "create dir failed: %+v", err)
				return fmt.Errorf("create dir failed")
			}
		}
	}

	// 设置上下文变量
	r.SetCtxVar("workspace", workspace)
	r.SetCtxVar("uploadOutput", uploadOutput)
	r.SetCtxVar("bots", bots)

	glog.Debugf(r.Context(), "mkdir, workspace: %s, uploadOutput: %s, bots: %s",
		workspace, uploadOutput, bots)

	return nil
}

// 设置语言
func (s *sMiddleware) setLang(r *ghttp.Request) {
	lang := ""

	// 从HTTP头获取Accept-Language
	acceptLang := r.GetHeader("Accept-Language")

	// 简化的语言检测逻辑
	if strings.Contains(acceptLang, "zh-CN") {
		lang = "zh-CN"
	} else if strings.Contains(acceptLang, "zh") {
		lang = "zh-CN" // 所有中文变体使用简体中文
	} else {
		lang = "en-US" // 默认值
	}

	r.SetCtxVar("lang", lang)
}

// TraceID use 'Trace-Id' from client request header by default.
func (s *sMiddleware) Prepare(r *ghttp.Request) {
	// 1. 准备工作目录
	if err := s.mkdir(r); err != nil {
		// 返回错误信息
		r.Response.WriteStatus(500)
		r.Response.WriteJson(g.Map{
			"code":    500,
			"message": "create dir failed",
		})
		r.ExitAll()
	}

	// 2. 设置语言环境
	s.setLang(r)

	r.Middleware.Next()
}
