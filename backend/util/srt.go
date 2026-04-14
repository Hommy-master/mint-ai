package util

import (
	"strings"
)

// 按行分割字符串并去除首尾的不可见字符
func ParseLines(input string) []string {
	// 按行分割字符串
	lines := strings.Split(input, "\n")

	// 创建一个切片来存储处理后的行
	var result []string

	// 遍历每一行，去除首尾的不可见字符
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" { // 如果处理后的行不为空，则添加到结果中
			result = append(result, trimmedLine)
		}
	}

	return result
}

func WriteTextToSRT(text, srtFile string) error {
	arr := ParseLines(text)
	return WriteLinesToFile(arr, srtFile)
}
