package wxgate

import (
	"context"
	"cozeos/internal/consts"
	"cozeos/internal/model"
	"cozeos/internal/pkg/cache"
	"cozeos/internal/pkg/db"
	"cozeos/internal/types"
	"fmt"
	"time"
)

// 获取随机验证码，返回一张图片，可以用HTML的<img>标签显示
func Scan(ctx context.Context, wxGate *types.WXGate) error {
	if wxGate.Ticket == "" {
		return nil
	}

	openID := wxGate.FromUserName
	_, err := db.CreateUserEx(ctx, &model.User{
		WeChatID:    openID,
		Name:        "jcaigc-" + openID,
		Phone:       "jcaigc-" + openID,
		VIPLevel:    0,
		VIPExpireAt: time.Now(),
	})
	if err != nil {
		return err
	}

	key := fmt.Sprintf(consts.TicketKeyFormat, wxGate.Ticket)
	cache.GetInstance().Set(key, openID, time.Duration(consts.WeChatQRCodeExpire)*time.Second)
	return nil
}
