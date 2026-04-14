package model

import "gorm.io/gorm"

// 充值消费记录表 - 这里表示的数值是以本站的积分为单位，并非实际支付全额（当前积分与实际支付金额等价）
type PluginBalanceLog struct {
	gorm.Model         // 包含默认字段：ID(流水号), CreatedAt, UpdatedAt, DeletedAt
	UserID      uint   `gorm:"column:user_id;index;foreignKey:ID;references:users"` // 用户ID (外键关联users表)
	User        User   `gorm:"foreignKey:UserID"`                                   // 关联用户模型（可选，便于预加载）
	Points      int64  `gorm:"column:points;type:bigint"`                           // 积分变动情况，正值表示积分充值，负值表示积分消费
	OrderNO     string `gorm:"column:order_no;type:varchar(64);index"`              // 订单号(唯一业务ID)，仅充值的时候不为空，其它情况这个值为空
	Description string `gorm:"column:description;type:varchar(128)"`                // 变动描述
	Ext         string `gorm:"column:ext;type:varchar(1024)"`                       // 扩展信息
}

func (p *PluginBalanceLog) TableName() string {
	return "plugin_balance_logs" // 指定表名
}
