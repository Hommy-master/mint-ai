package types

import "time"

type FileInfo struct {
	OriginalName string `json:"originalName" dc:"原始文件名"`
	FileSize     int64  `json:"fileSize" dc:"文件大小"`
	FileFormat   string `json:"fileFormat" dc:"文件格式"`
	URL          string `json:"url" dc:"文件访问URL"`
}

type Upload struct {
	ID           uint      `json:"id" dc:"文件ID"`
	UserID       uint      `json:"userID" dc:"用户ID"`
	OriginalName string    `json:"originalName" dc:"原始文件名"`
	DownloadURL  string    `json:"downloadURL" dc:"下载URL"`
	FileFormat   string    `json:"fileFormat" dc:"文件格式"`
	FileSize     uint64    `json:"fileSize" dc:"文件大小，单位：字节"`
	UploadTime   time.Time `json:"uploadTime" dc:"上传时间"`
	Remarks      string    `json:"remarks" dc:"备注"`
}
