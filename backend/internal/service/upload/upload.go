package upload

import (
	"context"
	"cozeos/internal/consts"
	"cozeos/internal/errcode"
	"cozeos/internal/model"
	"cozeos/internal/pkg/db"
	"cozeos/internal/pkg/helper"
	"cozeos/internal/pkg/i"
	"cozeos/internal/types"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

const (
	maxUploadSize = 100 * 1024 * 1024 // 最大上传100MB，与服务器配置保持一致
	price         = float64(0.08)
)

// 支持的文件类型
var supportedVideoTypes = map[string]bool{
	".mp4":  true,
	".mkv":  true,
	".avi":  true,
	".mov":  true,
	".wmv":  true,
	".flv":  true,
	".webm": true,
}

var supportedAudioTypes = map[string]bool{
	".mp3":  true,
	".wav":  true,
	".flac": true,
	".aac":  true,
	".ogg":  true,
	".wma":  true,
}

func Upload(ctx context.Context) (*types.FileInfo, error) {
	r := g.RequestFromCtx(ctx)
	uploadOutput := ctx.Value("uploadOutput").(string)

	// 0. 检查是否有足够的积分
	if _, err := db.CheckUserPoints(ctx, price); err != nil {
		glog.Warningf(ctx, "Upload check points failed, err: %+v", err)
		return nil, i.T(ctx, errcode.CodeInsufficientPoints)
	}

	// 1. 设置最大内存，超出部分将写入磁盘临时文件
	r.ParseMultipartForm(maxUploadSize)
	defer r.MultipartForm.RemoveAll()

	// 2. 获取上传的文件
	file, handler, err := r.FormFile("file")
	if err != nil {
		glog.Errorf(ctx, "Upload file not found, err: %+v", err)
		return nil, i.T(ctx, errcode.CodeUploadFileNotFound)
	}
	defer file.Close()

	// 3. 获取文件扩展名并验证
	ext := strings.ToLower(filepath.Ext(handler.Filename))
	if !isSupportedFormat(ext) {
		glog.Warningf(ctx, "Upload file format not supported, filename: %s", handler.Filename)
		return nil, i.T(ctx, errcode.CodeUploadFileFormatNotSupported)
	}

	// 4. 生成唯一文件名
	uniqueFilename := generateUniqueFilename(ext)
	dstPath := filepath.Join(uploadOutput, uniqueFilename)

	// 5. 创建目标文件
	dst, err := os.Create(dstPath)
	if err != nil {
		glog.Errorf(ctx, "Upload create file failed, err: %+v", err)
		return nil, i.T(ctx, errcode.CodeInternalServerError)
	}
	defer dst.Close()

	// 5. 复制文件内容
	if _, err := dst.ReadFrom(file); err != nil {
		glog.Errorf(ctx, "Upload copy file failed, err: %+v", err)
		return nil, i.T(ctx, errcode.CodeInternalServerError)
	}

	downloadURL := helper.GenerateDownloadURL(dstPath)

	// 6. 记录文件上传信息
	id := ctx.Value("id").(uint)
	glog.Infof(ctx, "user id: %d", id)
	upload := &model.Upload{
		UserID:       id,
		OriginalName: handler.Filename,
		DownloadURL:  downloadURL,
		FileFormat:   getFileFormat(ext),
		FileSize:     uint64(handler.Size),
		UploadTime:   time.Now(),
	}
	if err := db.NewDB().Create(upload).Error; err != nil {
		glog.Errorf(ctx, "Upload record failed, err: %+v", err)
		return nil, i.T(ctx, errcode.CodeInternalServerError)
	}

	// 7. 减少用户积分
	if err := db.UpdateUserPoints(ctx, id, -price, consts.FileUploadDesc, ""); err != nil {
		glog.Warningf(ctx, "Upload update points failed, err: %+v", err)
	}

	return &types.FileInfo{
		OriginalName: handler.Filename,
		FileSize:     handler.Size,
		FileFormat:   getFileFormat(ext),
		URL:          downloadURL,
	}, nil
}

func SupportedFormats(ctx context.Context) ([]string, []string, error) {
	return getSupportedVideoFormats(), getSupportedAudioFormats(), nil
}

// 查询文件上传记录
func QueryUpload(ctx context.Context, userID uint) ([]*types.Upload, error) {
	var records []*model.Upload
	if err := db.NewDB().Where("user_id = ?", userID).
		Order("upload_time DESC").Limit(1000).Find(&records).Error; err != nil {
		glog.Errorf(ctx, "QueryUploadRecords failed, err: %+v", err)
		return nil, i.T(ctx, errcode.CodeInternalServerError)
	}

	res := make([]*types.Upload, 0, len(records))
	for _, record := range records {
		res = append(res, helper.ModelUploadToTypesUpload(record))
	}

	return res, nil
}

// CleanOldUploadRecords 清理7天前的上传记录
func CleanOldUploadRecords(ctx context.Context) error {
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	result := db.NewDB().Where("upload_time < ?", sevenDaysAgo).Delete(&model.Upload{})
	if result.Error != nil {
		glog.Errorf(ctx, "CleanOldUploadRecords failed, err: %+v", result.Error)
		return result.Error
	}
	glog.Infof(ctx, "CleanOldUploadRecords success, deleted %d records", result.RowsAffected)
	return nil
}

// 检查文件格式是否支持
func isSupportedFormat(ext string) bool {
	return supportedVideoTypes[ext] || supportedAudioTypes[ext]
}

// 获取文件类型
func getFileFormat(ext string) string {
	if supportedVideoTypes[ext] {
		return "video"
	} else if supportedAudioTypes[ext] {
		return "audio"
	}
	return "unknown"
}

// 生成唯一文件名
func generateUniqueFilename(ext string) string {
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s%s", timestamp, ext)
}

// 获取支持的视频格式列表
func getSupportedVideoFormats() []string {
	formats := make([]string, 0, len(supportedVideoTypes))
	for ext := range supportedVideoTypes {
		formats = append(formats, strings.TrimPrefix(ext, "."))
	}
	return formats
}

// 获取支持的音频格式列表
func getSupportedAudioFormats() []string {
	formats := make([]string, 0, len(supportedAudioTypes))
	for ext := range supportedAudioTypes {
		formats = append(formats, strings.TrimPrefix(ext, "."))
	}
	return formats
}
