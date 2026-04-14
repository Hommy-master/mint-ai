package helper

import (
	"context"
	"testing"
)

func TestSendPaymentSuccessEmail(t *testing.T) {
	ctx := context.Background()
	err := SendPaymentSuccessEmail(ctx, 688001, "WEPAY20250709101030XDXDD", 100.0, "WEPAY20250709101030XDXDD", "2021-01-01 00:00:00")
	if err != nil {
		t.Errorf("SendPaymentSuccessEmail failed: %v", err)
	}
}

// TestSendBotCustomEmail 测试智能体定制通知邮件发送功能
func TestSendBotCustomEmail(t *testing.T) {
	ctx := context.Background()
	err := SendBotCustomEmail(ctx, "phone", "12345678901", "https://example.com/video.mp4", "¥1,500 - ¥3,000")
	if err != nil {
		t.Errorf("SendBotCustomEmail failed: %v", err)
	}
}
