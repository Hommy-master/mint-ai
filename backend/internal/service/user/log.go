package user

import (
	"context"
	"cozeos/internal/errcode"
	"cozeos/internal/model"
	"cozeos/internal/pkg/db"
	"cozeos/internal/pkg/i"
	"cozeos/internal/types"

	"github.com/gogf/gf/v2/os/glog"
)

// 查询积分充值/消费记录
func BalanceLog(ctx context.Context) ([]types.BalanceLog, error) {
	id := ctx.Value("id").(uint)

	// 1. 查询积充值、消费记录
	logs := []model.PluginBalanceLog{}
	err := db.NewDB().Model(&model.PluginBalanceLog{}).
		Where("user_id = ?", id).
		Order("created_at DESC").
		Limit(1000).
		Find(&logs).Error
	if err != nil {
		glog.Warningf(ctx, "query balance log failed, error: %v", err)
		return nil, i.T(ctx, errcode.CodeInternalServerError)
	}

	// 2. 数据转换
	res := []types.BalanceLog{}
	for _, log := range logs {
		res = append(res, types.BalanceLog{
			ID:          log.ID,
			OrderNO:     log.OrderNO,
			Points:      log.Points,
			Description: log.Description,
			OptAt:       log.CreatedAt,
		})
	}

	return res, nil
}
