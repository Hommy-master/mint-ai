package types

// 微信公众号网关回调结构体
type WXGate struct {
	EventType    string `json:"event_type" v:"required|in:subscribe,unsubscribe,scan" dc:"事件类型: 关注、取消关注、扫码"`
	OpenID       string `json:"open_id" v:"length:0,256" dc:"用户OpenID"`
	FromUserName string `json:"from_user_name" v:"required|length:1,256" dc:"发送者账号"`
	ToUserName   string `json:"to_user_name" v:"length:0,256" dc:"接收者账号"`
	Ticket       string `json:"ticket" v:"required|length:1,256" dc:"二维码票据（扫码事件）"`
	EventKey     string `json:"event_key" v:"required|length:1,128" dc:"事件键（扫码事件）"`
	CreateTime   int64  `json:"create_time" v:"required|min:1" dc:"消息创建时间戳"`
}
