package util

import (
	"fmt"
	"os"
)

func AppendToFile(filePath string, content string) error {
	// 修改打开文件的标志位，添加 os.O_CREATE 以在文件不存在时创建文件
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入内容
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func GetFileSize(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, fmt.Errorf("get file %s size failed: %v", filePath, err)
	}
	return info.Size(), nil
}

func WriteLinesToFile(lines []string, filePath string) error {
	// 打开文件以写入模式
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("create file %s failed, err: %w", filePath, err)
	}
	defer file.Close()

	// 遍历每一行，逐行写入文件
	for _, line := range lines {
		_, err := file.WriteString(line + "\n") // 添加换行符
		if err != nil {
			return fmt.Errorf("write string failed, err: %w", err)
		}
	}

	return nil
}
