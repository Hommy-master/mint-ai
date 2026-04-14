package upload

import (
	"context"
	v1 "cozeos/api/upload/v1"
	"cozeos/internal/service/upload"

	"github.com/gogf/gf/v2/os/glog"
)

func (c *ControllerV1) Upload(ctx context.Context, req *v1.UploadReq) (*v1.UploadRes, error) {
	glog.Debugf(ctx, "Upload req: %+v", *req)

	fileInfo, err := upload.Upload(ctx)
	if err != nil {
		glog.Warningf(ctx, "Upload err: %+v", err)
		return nil, err
	}

	return &v1.UploadRes{FileInfo: *fileInfo}, nil
}

func (c *ControllerV1) UploadSupportFormats(ctx context.Context, req *v1.UploadSupportFormatsReq) (*v1.UploadSupportFormatsRes, error) {
	glog.Debugf(ctx, "UploadSupportFormats req: %+v", *req)

	videoFormats, audioFormats, err := upload.SupportedFormats(ctx)
	if err != nil {
		glog.Warningf(ctx, "UploadSupportFormats err: %+v", err)
		return nil, err
	}

	return &v1.UploadSupportFormatsRes{VideoFormats: videoFormats, AudioFormats: audioFormats}, nil
}

func (c *ControllerV1) UploadRecords(ctx context.Context, req *v1.UploadRecordsReq) (*v1.UploadRecordsRes, error) {
	glog.Debugf(ctx, "UploadRecords req: %+v", *req)

	records, err := upload.QueryUpload(ctx, req.ID)
	if err != nil {
		glog.Warningf(ctx, "UploadRecords err: %+v", err)
		return nil, err
	}

	return &v1.UploadRecordsRes{Records: records}, nil
}
