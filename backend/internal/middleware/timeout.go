package middleware

import (
	"context"
	"cozeos/internal/errcode"
	"errors"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

const (
	timeout = 178 * time.Second // 设置HTTP请求超时时间
)

func (s *sMiddleware) TimeoutHandler(r *ghttp.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	r.SetCtx(ctx)

	// 使用缓冲通道（容量1）
	done := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1) // 添加panic通道

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		r.Middleware.Next()

		// 安全发送（因通道有缓冲不会阻塞）
		select {
		case done <- struct{}{}:
		default:
			// 超时后此处可能无接收者，但不会阻塞
		}
	}()

	select {
	case <-done:
		// 正常完成
	case p := <-panicChan:
		// 处理panic
		panic(p)
	case <-ctx.Done():
		// 超时处理
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			r.Response.ClearBuffer()
			r.Response.WriteJsonExit(response{
				Code:    errcode.CodeTimeout,
				Message: errcode.LocalizedMessage(errcode.CodeTimeout, ctx.Value("lang").(string)),
				TraceID: gctx.CtxId(r.Context()),
				Data:    nil,
			})

			// 关键修复：终止后续处理
			r.ExitAll()
		}
	}
}
