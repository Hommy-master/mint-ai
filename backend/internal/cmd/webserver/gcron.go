package webserver

import (
	"context"
	"cozeos/internal/service/upload"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/glog"
)

func initGCron(ctx context.Context) {
	// 每分钟执行一次的任务
	gcron.Add(ctx, "0 * * * * *", func(ctx context.Context) {

	})

	// 每天执行一次的任务
	gcron.Add(ctx, "0 0 0 * * *", func(ctx context.Context) {
		glog.Info(ctx, "daily task exec")

		// 清理7天前的上传记录
		upload.CleanOldUploadRecords(ctx)
	})
}
