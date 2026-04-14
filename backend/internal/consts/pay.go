package consts

import "time"

type PayMethod string // 支付方式

const (
	PayMethodWeChat PayMethod = "wechat"
	PayMethodAliPay PayMethod = "alipay"
)

type ProductID string // 商品

const (
	// 插件相关
	ProductPlugin20   ProductID = "PLUGIN_20"
	ProductPlugin50   ProductID = "PLUGIN_50"
	ProductPlugin100  ProductID = "PLUGIN_100"
	ProductPlugin200  ProductID = "PLUGIN_200"
	ProductPlugin500  ProductID = "PLUGIN_500"
	ProductPlugin1000 ProductID = "PLUGIN_1000"
	ProductCustom     ProductID = "CUSTOM" // 自定义充值

	// 会员充值
	ProductVipYear   ProductID = "VIP_YEAR"   // VIP年卡
	ProductSVipMonth ProductID = "SVIP_MONTH" // SVIP月卡
	ProductSVipYear  ProductID = "SVIP_YEAR"  // SVIP年卡

)

const (
	PayTimeout = 5 * time.Minute
)

type OrderStatus string // 支付状态

const (
	OrderStatusNotPay OrderStatus = "NOTPAY" // 未支付
	OrderStatusPayed  OrderStatus = "PAYED"  // 已支付
	OrderStatusCancel OrderStatus = "CANCEL" // 已取消
	OrderStatusRefund OrderStatus = "REFUND" // 已退款
	OrderStatusExpire OrderStatus = "EXPIRE" // 已过期
	OrderStatusFailed OrderStatus = "FAILED" // 支付失败
)

// 充值/消费描述信息
const (
	PluginRechargeDesc      = "插件充值"
	PluginRechargeBonusDesc = "插件充值赠送"
	NewUserRewardDesc       = "新人奖励"
	VideoPluginCallDesc     = "视频插件调用"
	VideoTransitionDesc     = "视频转场"
	FileUploadDesc          = "文件上传"
	VideoTrimCallDesc       = "视频裁剪调用"
	VideoRescaleCallDesc    = "视频缩放调用"
)
