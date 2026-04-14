package wxgate

import (
	"context"
	v1 "cozeos/api/wxgate/v1"
)

type IWXGateV1 interface {
	// 微信扫码
	Scan(ctx context.Context, req *v1.ScanReq) (res *v1.ScanRes, err error)
}
