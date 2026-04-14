package middleware

import (
	"cozeos/internal/config"
	"cozeos/internal/errcode"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gvalid"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	TraceID string      `json:"traceid"`
	Data    interface{} `json:"data"`
}

func (s *sMiddleware) isNeedHandle(r *ghttp.Request) bool {
	excludedUrls := []string{
		// 文档相关
		config.GetConfig().Server.SwaggerPath,
		config.GetConfig().Server.OpenApiPath,
		// 微信公众号对接，不做处理
		"/openapi/v1/wechat/officialaccount",
		"/openapi/v1/wechat/qrcode",
		// 微信支付 - 支付回调
		"/openapi/v1/user/pay/wechat-notify",
		// 流式接口，不需要在这里处理返回值
		"/stream/",
	}

	for _, url := range excludedUrls {
		if strings.HasPrefix(r.Request.URL.Path, url) {
			return false
		}
	}

	// 特殊处理
	if r.Request.URL.Path == "/openapi/v1/user/auth/captcha" { // 验证码
		if r.Request.Method == "GET" { // GET请求，不需要处理
			return false
		}
	}

	return true
}

// ResponseHandler custom response format.
func (s *sMiddleware) ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	// 过滤掉不需要处理的请求
	if !s.isNeedHandle(r) {
		glog.Infof(r.Context(), "No need to handle: %s", r.Request.URL.Path)
		return
	}

	// Clean exist response info.
	// i.e.:
	//	 1) gf exception recovered info
	//	 2) gf 404 Not Found content
	r.Response.ClearBuffer()

	var (
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code int
		msg  string
		lang = r.Context().Value("lang").(string)
	)
	// 处理不同错误类型 [6,7,8](@ref)
	switch {
	case err == nil:
		// 无错误时根据HTTP状态码设置业务码
		switch r.Response.Status {
		case 200:
			code = errcode.CodeSuccess
			msg = errcode.LocalizedMessage(errcode.CodeSuccess, lang)
		case 401:
			// 修改HTTP状态码
			r.Response.Status = 200
			code = errcode.CodeUnauthorized
			msg = errcode.LocalizedMessage(errcode.CodeUnauthorized, lang)
		case 403:
			code = errcode.CodeNotPermission
			msg = errcode.LocalizedMessage(errcode.CodeNotPermission, lang)
		case 500:
			code = errcode.CodeInternalServerError
			msg = errcode.LocalizedMessage(errcode.CodeInternalServerError, lang)
		default:
			code = errcode.CodeUnknown
			msg = errcode.LocalizedMessage(errcode.CodeUnknown, lang)
		}

	default:
		// 修复：使用类型断言代替IsValidationError [6,8](@ref)
		if validationErr, ok := err.(gvalid.Error); ok {
			code = errcode.CodeInvalidParameter
			// 获取第一条验证错误信息 [7,8](@ref)
			msg = errcode.LocalizedMessage(errcode.CodeInvalidParameter, lang) + ": " + gerror.Current(validationErr).Error()
		} else {
			tmpCode := gerror.Code(err)
			if tmpCode == nil {
				tmpCode = gcode.CodeInternalError
			}
			code = tmpCode.Code()
			msg = err.Error()
		}
	}

	r.Response.WriteJsonExit(response{
		Code:    code,
		Message: msg,
		TraceID: gctx.CtxId(r.Context()),
		Data:    res,
	})
}
