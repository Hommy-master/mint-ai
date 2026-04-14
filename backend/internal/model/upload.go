package model

import (
	"time"

	"gorm.io/gorm"
)

// Upload 文件上传记录表结构
type Upload struct {
	gorm.Model             // 包含默认字段：ID, CreatedAt, UpdatedAt, DeletedAt
	UserID       uint      `gorm:"column:user_id;not_null;index"` // 用户ID外键
	OriginalName string    `gorm:"column:original_name;type:varchar(256);not_null"`
	DownloadURL  string    `gorm:"column:download_url;type:varchar(256);not_null"`
	FileFormat   string    `gorm:"column:file_format;type:varchar(16);not_null"`
	FileSize     uint64    `gorm:"column:file_size;not_null"`
	UploadTime   time.Time `gorm:"column:upload_time;not_null"`
	Remarks      string    `gorm:"column:remarks;type:varchar(128)"`
	Ext          string    `gorm:"column:ext;type:varchar(1024);not_null"`
	User         User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"` // 关联用户
}

func (u *Upload) TableName() string {
	return "uploads" // 指定表名
}
