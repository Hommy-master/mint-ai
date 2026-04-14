package util

import (
	"fmt"
	"strconv"
)

func HexToRGB(hex string) (uint8, uint8, uint8, error) {
	// 去除可能的前缀并确保长度
	if len(hex) == 7 && hex[0] == '#' {
		hex = hex[1:]
	}
	if len(hex) != 6 {
		return 0, 0, 0, fmt.Errorf("无效的十六进制颜色格式")
	}

	// 解析R、G、B分量
	parseHex := func(s string) (int, error) {
		value, err := strconv.ParseInt(s, 16, 32)
		if err != nil {
			return 0, fmt.Errorf("无效的十六进制数字: %s", s)
		}
		return int(value), nil
	}

	r, err := parseHex(hex[0:2])
	if err != nil {
		return 0, 0, 0, err
	}

	g, err := parseHex(hex[2:4])
	if err != nil {
		return 0, 0, 0, err
	}

	b, err := parseHex(hex[4:6])
	if err != nil {
		return 0, 0, 0, err
	}

	return uint8(r), uint8(g), uint8(b), nil
}
