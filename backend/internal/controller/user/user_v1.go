package user

import (
	"context"
	v1 "cozeos/api/user/v1"
	"cozeos/internal/pkg/jwt"
	"cozeos/internal/service/user"

	"github.com/gogf/gf/v2/os/glog"
)

// 获取随机验证码，返回一张图片，可以用HTML的<img>标签显示
func (c *ControllerV1) AuthCaptcha(ctx context.Context, req *v1.AuthCaptchaReq) (*v1.AuthCaptchaRes, error) {
	glog.Debugf(ctx, "AuthCaptcha req: %+v", *req)

	// 生成验证码图片
	err := user.AuthCaptcha(ctx, req.Phone)
	if err != nil {
		glog.Warningf(ctx, "AuthCaptcha error: %v", err)
		return nil, err
	}

	return nil, nil
}

// 验证验证码
func (c *ControllerV1) AuthCaptchaVerify(ctx context.Context, req *v1.AuthCaptchaVerifyReq) (*v1.AuthCaptchaVerifyRes, error) {
	glog.Debugf(ctx, "AuthCaptchaVerify req: %+v", *req)

	// 验证验证码
	err := user.AuthCaptchaVerify(ctx, req.Phone, req.Code)
	if err != nil {
		glog.Warningf(ctx, "AuthCaptchaVerify error: %v", err)
		return nil, err
	}

	return nil, nil
}

// 发送短信验证码
func (c *ControllerV1) AuthSms(ctx context.Context, req *v1.AuthSmsReq) (*v1.AuthSmsRes, error) {
	return nil, nil
}

// 验证短信验证码
func (c *ControllerV1) AuthSmsVerify(ctx context.Context, req *v1.AuthSmsVerifyReq) (*v1.AuthSmsVerifyRes, error) {
	return nil, nil
}

func (c *ControllerV1) AuthLogin(ctx context.Context, req *v1.AuthLoginReq) (*v1.AuthLoginRes, error) {
	glog.Debugf(ctx, "AuthLogin req: %+v", *req)

	// 登录
	u, err := user.AuthLogin(ctx, req.Phone, req.Pass)
	if err != nil {
		glog.Warningf(ctx, "AuthLogin error: %v", err)
		return nil, err
	}

	// 生成JWT令牌
	token, err := jwt.GenerateToken(u.ID, u.Phone, u.WeChatID)
	if err != nil {
		glog.Errorf(ctx, "generate jwt token failed: %+v", err)
		return nil, err
	}

	return &v1.AuthLoginRes{User: *u, Token: token}, nil
}

func (c *ControllerV1) Check(ctx context.Context, req *v1.CheckReq) (*v1.CheckRes, error) {
	glog.Debugf(ctx, "Check req: %+v", *req)

	// 检查用户是否存在
	exists, err := user.Check(ctx, req.Phone)
	if err != nil {
		glog.Warningf(ctx, "Check error: %v", err)
		return nil, err
	}

	return &v1.CheckRes{Exists: exists}, nil
}

// 获取二维码
func (c *ControllerV1) AuthQRCode(ctx context.Context, req *v1.AuthQRCodeReq) (*v1.AuthQRCodeRes, error) {
	glog.Debugf(ctx, "AuthQRCode req: %+v", *req)

	ticket, expire, err := user.AuthQRCode(ctx)
	if err != nil {
		glog.Warningf(ctx, "AuthQRCode error: %v", err)
		return nil, err
	}

	return &v1.AuthQRCodeRes{Ticket: ticket, Expire: expire}, nil
}

// 查询扫码状态
func (c *ControllerV1) AuthQRCodeStatus(ctx context.Context, req *v1.AuthQRCodeStatusReq) (*v1.AuthQRCodeStatusRes, error) {
	glog.Debugf(ctx, "AuthQRCodeStatus req: %+v", *req)

	u, err := user.AuthQRCodeStatus(ctx, req.SceneID, req.Ticket)
	if err != nil {
		glog.Infof(ctx, "AuthQRCodeStatus error: %v", err)
		return nil, err
	}

	// 生成JWT令牌
	token, err := jwt.GenerateToken(u.ID, u.Phone, u.WeChatID)
	if err != nil {
		glog.Errorf(ctx, "generate jwt token failed: %+v", err)
		return nil, err
	}

	return &v1.AuthQRCodeStatusRes{User: *u, Token: token}, nil
}

func (c *ControllerV1) QueryUser(ctx context.Context, req *v1.QueryUserReq) (*v1.QueryUserRes, error) {
	glog.Debugf(ctx, "QueryUser req: %+v", *req)

	u, err := user.QueryUser(ctx, req.ID)
	if err != nil {
		glog.Warningf(ctx, "QueryUser error: %v", err)
		return nil, err
	}

	return &v1.QueryUserRes{User: *u}, nil
}

func (c *ControllerV1) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (*v1.UpdateUserRes, error) {
	glog.Debugf(ctx, "UpdateUser req: %+v", *req)

	err := user.UpdateUser(ctx, req.ID, req.Name, req.APIKey)
	if err != nil {
		glog.Warningf(ctx, "UpdateUser error: %v", err)
		return nil, err
	}

	return &v1.UpdateUserRes{}, nil
}

// 微信支付 - 创建订单
func (c *ControllerV1) CreateOrder(ctx context.Context, req *v1.CreateOrderReq) (*v1.CreateOrderRes, error) {
	glog.Debugf(ctx, "CreateOrder req: %+v", *req)

	qrCodeURL, orderNO, expire, price, err := user.CreateOrder(ctx, req.PayMethod, req.ProductID, req.ProductValue)
	if err != nil {
		glog.Warningf(ctx, "CreateOrder error: %v", err)
		return nil, err
	}

	return &v1.CreateOrderRes{QRCodeURL: qrCodeURL, OrderNO: orderNO, Expire: expire, Price: price}, nil

}

// 微信支付 - 支付结果回调
func (c *ControllerV1) WeChatNotify(ctx context.Context, req *v1.WeChatNotifyReq) (*v1.WeChatNotifyRes, error) {
	glog.Debugf(ctx, "WeChatNotify req: %+v", *req)
	user.WeChatNotify(ctx)
	return &v1.WeChatNotifyRes{}, nil
}

// 微信支付 - 查询支付结果
func (c *ControllerV1) CheckPayment(ctx context.Context, req *v1.CheckPaymentReq) (*v1.CheckPaymentRes, error) {
	glog.Debugf(ctx, "CheckPayment req: %+v", *req)

	status, err := user.CheckPayment(ctx, req.OrderNO)
	if err != nil {
		glog.Warningf(ctx, "CheckPayment error: %v", err)
		return nil, err
	}

	return &v1.CheckPaymentRes{Status: status}, nil
}

func (c *ControllerV1) BalanceLog(ctx context.Context, req *v1.BalanceLogReq) (*v1.BalanceLogRes, error) {
	glog.Debugf(ctx, "BalanceLog req: %+v", *req)

	logs, err := user.BalanceLog(ctx)
	if err != nil {
		glog.Warningf(ctx, "BalanceLog error: %v", err)
		return nil, err
	}

	return &v1.BalanceLogRes{Logs: logs}, nil
}

func (c *ControllerV1) Query(ctx context.Context, req *v1.QueryReq) (res *v1.QueryRes, err error) {
	pt, err := user.Query(ctx, req.APIKey)
	if err != nil {
		glog.Warningf(ctx, "points.Query error: %v", err)
		return nil, err
	}

	return &v1.QueryRes{Points: pt}, nil
}

func (c *ControllerV1) Deduct(ctx context.Context, req *v1.DeductReq) (res *v1.DeductRes, err error) {
	err = user.Deduct(ctx, req.APIKey, req.Points, req.Desc)
	if err != nil {
		glog.Warningf(ctx, "points.Deduct error: %v", err)
		return nil, err
	}

	return &v1.DeductRes{}, nil
}
