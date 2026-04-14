package model

import (
	"time"

	"gorm.io/gorm"
)

// 核心订单表
type Order struct {
	gorm.Model              // 包含默认字段：ID, CreatedAt, UpdatedAt, DeletedAt
	UserID        uint      `gorm:"column:user_id;index;foreignKey:ID;references:users"` // 用户ID (外键关联users表)
	User          User      `gorm:"foreignKey:UserID"`                                   // 关联用户模型（可选，便于预加载）
	ProductID     string    `gorm:"column:product_id;type:varchar(64);index"`            // 产品ID
	OrderNO       string    `gorm:"column:order_no;type:varchar(64);uniqueIndex"`        // 订单号(唯一业务ID)
	Amount        float64   `gorm:"column:amount;type:decimal(10,2)"`                    // 订单总金额，单位：人民币
	PayMethod     string    `gorm:"column:pay_method;type:varchar(32)"`                  // 支付方式
	TransactionID string    `gorm:"column:transaction_id;type:varchar(64)"`              // 支付平台交易号
	PayedAt       time.Time `gorm:"column:payed_at;index"`                               // 支付时间
	OrderStatus   string    `gorm:"column:order_status;type:varchar(64);index"`          // 订单状态
	Payer         string    `gorm:"column:payer;type:varchar(64)"`                       // 支付人
	Ext           string    `gorm:"column:ext;type:varchar(1024)"`                       // 扩展信息
}

func (o *Order) TableName() string {
	return "orders" // 指定表名
}
