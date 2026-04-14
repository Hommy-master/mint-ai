package v1

import (
	"cozeos/internal/types"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取随机验证码，返回一张图片，可以用HTML的<img>标签显示
type AuthCaptchaReq struct {
	g.Meta `path:"/auth/captcha" method:"get" tags:"用户管理" summary:"获取随机验证码" dc:"生成随机验证码图片, 可以用HTML的<img>标签显示, 且必需带查询参数phone, 需要保证手机号码格式正确"`
	Phone  string `query:"phone" v:"required|phone" dc:"手机号"`
}

type AuthCaptchaRes struct {
}

// 验证验证码
type AuthCaptchaVerifyReq struct {
	g.Meta `path:"/auth/captcha/verify" method:"post" tags:"用户管理" summary:"验证验证码" dc:"验证码必需一次输入正确，否则需要重新刷新验证码"`
	Phone  string `json:"phone" v:"required|phone" dc:"手机号"`
	Code   string `json:"code" v:"required|length:4,4" dc:"验证码"`
}

type AuthCaptchaVerifyRes struct {
}

// 发送短信验证码
type AuthSmsReq struct {
	g.Meta `path:"/auth/sms" method:"post" tags:"用户管理" summary:"发送短信验证码" dc:"发送短信验证码"`
	Phone  string `json:"phone" v:"required|phone" dc:"手机号"`
}

type AuthSmsRes struct {
}

// 验证短信验证码
type AuthSmsVerifyReq struct {
	g.Meta `path:"/auth/sms/verify" method:"post" tags:"用户管理" summary:"验证短信验证码" dc:"验证短信验证码"`
	Phone  string `json:"phone" v:"required|phone" dc:"手机号"`
	Code   string `json:"code" v:"required|length:4,4" dc:"验证码"`
}

type AuthSmsVerifyRes struct {
}

type AuthLoginReq struct {
	g.Meta `path:"/auth/login" method:"post" tags:"用户管理" summary:"用户名密码登录" dc:"用户名密码登录，注册和登录是同一个接口"`
	Phone  string `json:"phone" v:"required|phone" dc:"手机号"`
	Pass   string `json:"pass" v:"required|length:6,32" dc:"密码, 长度6-32个字符"`
}

type AuthLoginRes struct {
	types.User
	Token string `json:"token" dc:"JWT令牌"`
}

type CheckReq struct {
	g.Meta `path:"/check" method:"get" tags:"用户管理" summary:"检查用户是否存在" dc:"检查用户是否存在"`
	Phone  string `query:"phone" v:"required|phone" dc:"手机号"`
}

type CheckRes struct {
	Exists bool `json:"exists" dc:"是否存在"`
}

type AuthQRCodeReq struct {
	g.Meta `path:"/auth/qrcode" method:"get" tags:"用户管理" summary:"获取二维码" dc:"获取服务号二维码，用于扫码登录"`
}

type AuthQRCodeRes struct {
	Ticket string `json:"ticket" dc:"二维码ticket"`
	Expire int    `json:"expire" dc:"二维码过期时间，单位：秒"`
}

type AuthQRCodeStatusReq struct {
	g.Meta  `path:"/auth/qrcode/status" method:"get" tags:"用户管理" summary:"获取扫码登录状态" dc:"获取用户扫码登录状态，需要做超时处理，根据上一步返回的二维码过期时间来处理"`
	SceneID string `query:"sceneID" v:"required|size:5" dc:"场景ID"`
	Ticket  string `query:"ticket" v:"required" dc:"二维码ticket"`
}

type AuthQRCodeStatusRes struct {
	types.User
	Token string `json:"token" dc:"JWT令牌"`
}

type QueryUserReq struct {
	g.Meta `path:"/" method:"get" tags:"用户管理" summary:"查询用户信息" dc:"根据用户ID查询用户信息"`
	ID     uint `query:"id" v:"required|min:1" dc:"用户ID"`
}

type QueryUserRes struct {
	types.User
}

type UpdateUserReq struct {
	g.Meta `path:"/" method:"post" tags:"用户管理" summary:"更新用户信息" dc:"根据用户ID更新用户信息"`
	ID     uint    `json:"id" v:"required|min:1" dc:"用户ID"`
	Name   *string `json:"name" v:"length:1,64" dc:"用户名"`
}

type UpdateUserRes struct {
}

// 微信支付 - 创建订单
type CreateOrderReq struct {
	g.Meta       `path:"/pay/create-order" method:"post" tags:"用户管理" summary:"创建支付订单" dc:"根据支付类型创建支付订单"`
	PayMethod    string `json:"payMethod" v:"required|size:6" dc:"支付方式，只支持: wechat, 暂时不支持: alipay"`
	ProductID    string `json:"productID" v:"required|length:6,32" dc:"产品ID, 如: PLUGIN_20、PLUGIN_50、PLUGIN_100、PLUGIN_200、PLUGIN_500、CUSTOM"`
	ProductValue int    `json:"productValue" v:"required-if:productID,CUSTOM|min:1" dc:"产品价格, 仅插件充值自定义金额时用到，单位：元"`
}

type CreateOrderRes struct {
	QRCodeURL string `json:"qrCodeURL" dc:"二维码URL"`
	OrderNO   string `json:"orderNO" dc:"订单号"`
	Expire    int    `json:"expire" dc:"订单过期时间, 默认300秒过期, 单位：秒"`
	Price     int    `json:"price" dc:"订单金额, 单位：元"`
}

// 微信支付 - 支付结果回调
type WeChatNotifyReq struct {
	g.Meta `path:"/pay/wechat-notify" method:"post" tags:"用户管理" summary:"支付结果回调" dc:"支付结果回调"`
}

type WeChatNotifyRes struct {
}

// 微信支付 - 查询支付结果
type CheckPaymentReq struct {
	g.Meta  `path:"/pay/check-payment" method:"get" tags:"用户管理" summary:"查询支付结果" dc:"查询支付结果"`
	OrderNO string `query:"orderNO" v:"required|size:25" dc:"订单号"`
}

type CheckPaymentRes struct {
	Status string `json:"status" dc:"支付状态, 取值: NOTPAY(未支付), PAYED(已支付), CANCEL(已取消), REFUND(已退款), EXPIRE(已过期)"`
}

type BalanceLogReq struct {
	g.Meta `path:"/balance/log" method:"get" tags:"用户管理" summary:"查询充值消费记录" dc:"查询充值/消费记录信息"`
}

type BalanceLogRes struct {
	Logs []types.BalanceLog `json:"logs" dc:"交易日志记录"`
}

// 查询用户积分
type QueryReq struct {
	g.Meta `path:"/points" method:"get" tags:"查询用户积分" summary:"查询用户积分" dc:"基于JWT登录态查询当前用户积分"`
}

type QueryRes struct {
	Points int64 `json:"points" dc:"用户积分"`
}

type DeductReq struct {
	g.Meta `path:"/points/deduct" method:"post" tags:"扣除用户积分" summary:"扣除用户积分" dc:"基于JWT登录态扣除当前用户积分"`
	Points int64  `json:"points" v:"required|integer|between:1,1000" dc:"当前消耗积分，必需是一个正整数"`
	Desc   string `json:"desc" v:"required|length:1,128" dc:"当前消耗积分描述，是什么场景下消耗积分"`
}

type DeductRes struct {
}
