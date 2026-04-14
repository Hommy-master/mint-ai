package wxgate

import (
	"context"
	v1 "cozeos/api/wxgate/v1"
	"cozeos/internal/service/wxgate"

	"github.com/gogf/gf/v2/os/glog"
)

// 获取随机验证码，返回一张图片，可以用HTML的<img>标签显示
func (c *ControllerV1) Scan(ctx context.Context, req *v1.ScanReq) (*v1.ScanRes, error) {
	glog.Debugf(ctx, "Scan req: %+v", req.WXGate)

	// 生成验证码图片
	err := wxgate.Scan(ctx, &req.WXGate)
	if err != nil {
		glog.Warningf(ctx, "AuthCaptcha error: %v", err)
		return nil, err
	}

	return nil, nil
}
