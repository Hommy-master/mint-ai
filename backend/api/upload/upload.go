package upload

import (
	"context"
	v1 "cozeos/api/upload/v1"
)

type IUploadV1 interface {
	// 文件上传接口
	Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error)
	// 文件上传接口支持的文件格式
	UploadSupportFormats(ctx context.Context, req *v1.UploadSupportFormatsReq) (res *v1.UploadSupportFormatsRes, err error)
	// 文件上传记录
	UploadRecords(ctx context.Context, req *v1.UploadRecordsReq) (res *v1.UploadRecordsRes, err error)
}
