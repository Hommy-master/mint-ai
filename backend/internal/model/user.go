package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	AUTO_INCREMENT = 688000 // 自增起始值
)

// 用户表
type User struct {
	gorm.Model            // 包含默认字段：ID, CreatedAt, UpdatedAt, DeletedAt
	Name        string    `gorm:"column:name;type:varchar(64);not_null;uniqueIndex"`      // 用户名
	Phone       string    `gorm:"column:phone;type:varchar(64);not_null;uniqueIndex"`     // 手机号码
	Pass        string    `gorm:"column:pass;type:varchar(64);default:null"`              // 密码
	WeChatID    string    `gorm:"column:wechat_id;type:varchar(64);not_null;uniqueIndex"` // 关联微信 ID
	Email       string    `gorm:"column:email;type:varchar(256);default:null"`            // 邮箱
	Role        string    `gorm:"column:role;type:varchar(32);default:'normal'"`          // 角色（normal:普通用户 creator:创作者 admin:管理员）
	Points      int64     `gorm:"column:points;type:bigint;default:1"`                    // 积分余额（整数）
	VIPLevel    int       `gorm:"column:vip_level;type:int;default:0;index"`              // VIP等级（0:非VIP 1:VIP_YEAR 2:SVIP_MONTH 3:SVIP_YEAR）
	VIPExpireAt time.Time `gorm:"column:vip_expire_at;index;default:0001-01-01 00:00:00"` // VIP到期时间
	ReferrerID  uint      `gorm:"column:referrer_id;index;default:0"`                     // 推荐人ID，0表示没有推荐人（明确设置default:0）
	Ext         string    `gorm:"column:ext;type:varchar(1024)"`                          // 扩展字段，使用 JSON 存储
}

func (u *User) TableName() string {
	return "users" // 指定表名
}

/*
 1. 如何查询某个指定用户推荐了多少个用户？
 答：通过查询ReferrerID=指定用户ID的所有用户记录

 2. 如果查询当前用户推荐的用户消费了多少钱？
 答：先通过ReferrerID查询到所有推荐的用户，再通过用户ID查询消费记录

 3. 工作流商店-> 创建者中心 -> 收益统计
 答：先去掉这个功能
 **/
