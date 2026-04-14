package v1

import (
	"cozeos/internal/types"

	"github.com/gogf/gf/v2/frame/g"
)

// 文件上传接口
type UploadReq struct {
	g.Meta `path:"/" method:"post" tags:"文件上传" summary:"文件上传接口" dc:"文件上传接口"`
}

type UploadRes struct {
	types.FileInfo
}

// 文件上传接口支持的文件格式
type UploadSupportFormatsReq struct {
	g.Meta `path:"/supportformats" method:"get" tags:"文件上传" summary:"文件上传接口支持的文件格式" dc:"获取文件上传接口支持的文件格式"`
}

type UploadSupportFormatsRes struct {
	VideoFormats []string `json:"videoFormats" dc:"视频文件格式"`
	AudioFormats []string `json:"audioFormats" dc:"音频文件格式"`
}

type UploadRecordsReq struct {
	g.Meta `path:"/records" method:"get" tags:"文件上传" summary:"文件上传记录" dc:"获取文件上传记录"`
	ID     uint `query:"id" dc:"用户ID" v:"required"`
}

type UploadRecordsRes struct {
	Records []*types.Upload `json:"records" dc:"文件上传记录"`
}
