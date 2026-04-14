package middleware

import (
	"cozeos/internal/pkg/jwt"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func (s *sMiddleware) Auth(r *ghttp.Request) {
	// 定义排除路径（使用map提高查询效率）
	excludedPaths := map[string]bool{
		"/openapi/v1/user/auth/login":         true,
		"/openapi/v1/user/auth/qrcode":        true,
		"/openapi/v1/user/auth/qrcode/status": true,
		"/openapi/v1/user/pay/wechat-notify":  true,
		"/openapi/v1/user/points":             true, // 外部插件调用，根据apiKey查询用户积分，不进行认证
		"/openapi/v1/user/points/deduct":      true, // 外部插件调用，根据apiKey扣除用户积分，不进行认证
	}

	// 获取当前请求路径和方法
	currentPath := r.URL.Path
	currentMethod := r.Method
	glog.Infof(r.Context(), "current path: %s, method: %s", currentPath, currentMethod)

	// 检查当前路径是否需要排除认证
	// 对于 /openapi/v1/bots 路径，只有GET请求不需要认证
	if _, excluded := excludedPaths[currentPath]; excluded ||
		(currentPath == "/openapi/v1/bots" && currentMethod == "GET") ||
		(currentPath == "/openapi/v1/bots/top" && currentMethod == "GET") {
		r.Middleware.Next()
		return
	}

	// 从请求头中获取JWT令牌
	authHeader := r.Header.Get("Authorization")
	tokenString, err := jwt.ExtractTokenFromHeader(authHeader)
	if err != nil {
		glog.Warningf(r.Context(), "extract JWT token failed: %v, path: %s", err, currentPath)
		// 返回401状态码
		r.Response.WriteStatusExit(401, "Unauthorized")
		return
	}

	// 验证JWT令牌
	claims, err := jwt.ValidateToken(tokenString)
	if err != nil {
		glog.Warningf(r.Context(), "validate JWT token failed: %v, path: %s", err, currentPath)
		// 返回401状态码
		r.Response.WriteStatusExit(401, "Unauthorized")
		return
	}

	// 检查用户ID是否有效
	if claims.UserID < 1 {
		glog.Warningf(r.Context(), "user id not found in JWT claims, path: %s", currentPath)
		// 返回401状态码
		r.Response.WriteStatusExit(401, "Unauthorized")
		return
	}

	// 设置用户ID到上下文
	r.SetCtxVar("id", claims.UserID)

	// 用户已认证，继续后续流程
	r.Middleware.Next()
}
