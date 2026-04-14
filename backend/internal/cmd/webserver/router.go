package webserver

import (
	"context"
	"cozeos/internal/controller/upload"
	"cozeos/internal/controller/user"
	"cozeos/internal/controller/wxgate"
	"cozeos/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func wxgateRouter(ctx context.Context, s *ghttp.Server) {
	glog.Debugf(ctx, "set wxgate router")

	// 用户管理相关接口
	s.Group("/openapi/v1/wxgate", func(group *ghttp.RouterGroup) {
		group.Bind(wxgate.NewV1())
	})
}

func userRouter(ctx context.Context, s *ghttp.Server) {
	glog.Debugf(ctx, "set user router")

	// 用户管理相关接口
	s.Group("/openapi/v1/user", func(group *ghttp.RouterGroup) {
		m := middleware.New()
		group.Middleware(m.Auth)
		group.Bind(user.NewV1())
	})
}

func uploadRouter(ctx context.Context, s *ghttp.Server) {
	glog.Debugf(ctx, "set upload router")

	// 文件上传相关接口
	s.Group("/openapi/v1/upload", func(group *ghttp.RouterGroup) {
		m := middleware.New()
		group.Middleware(m.Auth)
		group.Bind(upload.NewV1())
	})
}

func initRouter(ctx context.Context, s *ghttp.Server) {
	// 全局中间件
	m := middleware.New()
	s.Use(
		m.TraceID,
		m.Prepare,
		m.TimeoutHandler,
		m.ResponseHandler,
	)

	// 微信公众号网关回调接口
	wxgateRouter(ctx, s)

	// 用户管理相关接口
	userRouter(ctx, s)

	// 文件上传相关接口
	uploadRouter(ctx, s)
}
