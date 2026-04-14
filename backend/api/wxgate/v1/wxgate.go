package v1

import (
	"cozeos/internal/types"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取随机验证码，返回一张图片，可以用HTML的<img>标签显示
type ScanReq struct {
	g.Meta `path:"/scan" method:"post" tags:"微信网关" summary:"微信扫码登录回调" dc:"用户扫码成功后，微信公众号网关回调"`
	types.WXGate
}

type ScanRes struct {
}
