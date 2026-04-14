package util

import (
	"fmt"
	"testing"
)

func TestHexToRGB(t *testing.T) {
	// 测试用例
	testCases := []string{
		"F0DAC5",  // 正常值
		"#F0DAC5", // 带#前缀
		"FF0000",  // 红色
		"00FF00",  // 绿色
		"0000FF",  // 蓝色
		"FFFFFF",  // 白色
		"000000",  // 黑色
	}

	for _, tc := range testCases {
		r, g, b, err := HexToRGB(tc)
		if err != nil {
			fmt.Printf("错误: %s -> %v\n", tc, err)
		} else {
			fmt.Printf("%-7s -> R:%3d, G:%3d, B:%3d\n", tc, r, g, b)
		}
	}
}
