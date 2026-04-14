package types

type Color struct {
	R uint8 `json:"r" v:"required|min:0|max:255" dc:"红色分量"`
	G uint8 `json:"g" v:"required|min:0|max:255" dc:"绿色分量"`
	B uint8 `json:"b" v:"required|min:0|max:255" dc:"蓝色分量"`
	A uint8 `json:"a" dc:"透明度分量"`
}
