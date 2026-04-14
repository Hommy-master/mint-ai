package db

import (
	"context"
	"cozeos/internal/model"

	"github.com/gogf/gf/v2/os/glog"
)

// 添加一条充值/消费记录
func AddBalanceLog(ctx context.Context, bl *model.PluginBalanceLog) error {
	err := NewDB().Model(&model.PluginBalanceLog{}).Create(bl).Error
	if err != nil {
		glog.Errorf(ctx, "AddBalanceLog failed: %v", err)
		return err
	}

	return nil
}
