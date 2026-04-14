package wxgate

import "cozeos/api/wxgate"

type ControllerV1 struct{}

func NewV1() wxgate.IWXGateV1 {
	return &ControllerV1{}
}
