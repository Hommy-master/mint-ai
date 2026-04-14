package types

import (
	"cozeos/internal/consts"
	"time"
)

type User struct {
	ID          uint      `json:"id" dc:"用户ID"`                                              // 主键ID，自增
	Name        string    `json:"name" dc:"用户名"`                                             // 用户名，最大长度为255个字符
	Email       string    `json:"email" dc:"邮箱"`                                             // 邮箱，最大长度为255个字符
	Phone       string    `json:"phone" dc:"手机号"`                                            // 手机号，最大长度为20个字符
	Points      float64   `json:"points" dc:"积分/资源点"`                                        // 单位：元
	WeChatID    string    `json:"wechatID" dc:"关联微信, 即微信的OpenID"`                            // 微信ID，最大长度为255个字符
	Role        string    `json:"role" dc:"角色，取值：normal: 普通用户, creator: 智能体创建者, admin: 管理员"` // 用户角色，主要用于控制哪些用户可以创建智能体
	VIPLevel    int       `json:"vipLevel" dc:"VIP等级 (0:非VIP 1:VIP1 2:VIP2) "`               // VIP等级（0:非VIP 1:VIP1 2:VIP2）
	VIPExpireAt time.Time `json:"vipExpireAt" dc:"VIP到期时间"`                                  // VIP到期时间
	ReferrerID  uint      `json:"referrerID" dc:"推荐人ID"`
	Ext         string    `json:"ext" dc:"扩展字段，使用 JSON 存储"` // 扩展字段，使用 JSON 存储
}

// 微信支付 - 通知
type WeChatPayNotify struct {
	ID           string `json:"id"`
	CreateTime   string `json:"create_time"`
	ResourceType string `json:"resource_type"`
	EventType    string `json:"event_type"`
	Summary      string `json:"summary"`
	Resource     struct {
		Algorithm      string `json:"algorithm"`
		Ciphertext     string `json:"ciphertext"`
		AssociatedData string `json:"associated_data"`
		Nonce          string `json:"nonce"`
	} `json:"resource"`
}

// DecryptedResource 解密后的资源数据
type DecryptedResource struct {
	AppID          string `json:"appid"`
	MchID          string `json:"mchid"`
	OutTradeNo     string `json:"out_trade_no"`
	TransactionID  string `json:"transaction_id"`
	TradeType      string `json:"trade_type"`
	TradeState     string `json:"trade_state"`
	TradeStateDesc string `json:"trade_state_desc"`
	BankType       string `json:"bank_type"`
	Attach         string `json:"attach"`
	SuccessTime    string `json:"success_time"`
	Amount         struct {
		Total    int64  `json:"total"`
		Currency string `json:"currency"`
	} `json:"amount"`
	Payer struct {
		OpenID string `json:"openid"`
	} `json:"payer"`
}

type OrderInfo struct {
	UserID      uint               `json:"userID"`
	ProductID   consts.ProductID   `json:"productID"`
	Price       int64              `json:"price"`
	OrderStatus consts.OrderStatus `json:"orderStatus"`
}

type BalanceLog struct {
	ID          uint      `json:"id" dc:"日志ID"`
	OrderNO     string    `json:"orderNO" dc:"订单号"`
	Points      float64   `json:"points" dc:"积分/资源点"`
	Description string    `json:"description" dc:"交易详情"`
	OptAt       time.Time `json:"optAt" dc:"操作时间"`
}
