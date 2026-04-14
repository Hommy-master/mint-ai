package user

import (
	"context"
	v1 "cozeos/api/user/v1"
)

type IUserV1 interface {
	// 获取随机验证码，返回一张图片，可以用HTML的<img>标签显示
	AuthCaptcha(ctx context.Context, req *v1.AuthCaptchaReq) (res *v1.AuthCaptchaRes, err error)
	// 验证验证码
	AuthCaptchaVerify(ctx context.Context, req *v1.AuthCaptchaVerifyReq) (res *v1.AuthCaptchaVerifyRes, err error)
	// 发送短信验证码
	AuthSms(ctx context.Context, req *v1.AuthSmsReq) (res *v1.AuthSmsRes, err error)
	// 验证短信验证码
	AuthSmsVerify(ctx context.Context, req *v1.AuthSmsVerifyReq) (res *v1.AuthSmsVerifyRes, err error)
	// 用户名密码登录
	AuthLogin(ctx context.Context, req *v1.AuthLoginReq) (res *v1.AuthLoginRes, err error)
	// 检查用户是否存在
	Check(ctx context.Context, req *v1.CheckReq) (res *v1.CheckRes, err error)
	// 获取二维码
	AuthQRCode(ctx context.Context, req *v1.AuthQRCodeReq) (res *v1.AuthQRCodeRes, err error)
	// 查询扫码状态 - 查询用户登录状态
	AuthQRCodeStatus(ctx context.Context, req *v1.AuthQRCodeStatusReq) (res *v1.AuthQRCodeStatusRes, err error)
	// 查询用户信息
	QueryUser(ctx context.Context, req *v1.QueryUserReq) (res *v1.QueryUserRes, err error)
	// 更新用户信息
	UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error)
	// 微信支付 - 创建订单
	CreateOrder(ctx context.Context, req *v1.CreateOrderReq) (res *v1.CreateOrderRes, err error)
	// 微信支付 - 支付结果回调
	WeChatNotify(ctx context.Context, req *v1.WeChatNotifyReq) (res *v1.WeChatNotifyRes, err error)
	// 微信支付 - 查询支付结果
	CheckPayment(ctx context.Context, req *v1.CheckPaymentReq) (res *v1.CheckPaymentRes, err error)
	// 查询充值消费记录
	BalanceLog(ctx context.Context, req *v1.BalanceLogReq) (res *v1.BalanceLogRes, err error)
	// 查询用户积分
	Query(ctx context.Context, req *v1.QueryReq) (res *v1.QueryRes, err error)
	// 扣除用户积分
	Deduct(ctx context.Context, req *v1.DeductReq) (res *v1.DeductRes, err error)
}
