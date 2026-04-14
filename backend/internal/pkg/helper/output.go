package helper

import (
	"cozeos/internal/config"
	"cozeos/internal/consts"
	"fmt"
	"os"
	"strings"
)

// 生成下载URL，参数示例：/app/output/image/2025-05-01/1746093311_4_0001.mp4
func GenerateDownloadURL(localPath string) string {
	newPath := strings.Replace(localPath, consts.WebRootPrefix, config.URL, 1)
	if strings.HasPrefix(newPath, config.URL) {
		return newPath
	} else {
		panic(fmt.Sprintf("invalid local path, localPath: %s", localPath)) // 这里可以抛出一个错误或者返回一个默认的URL，根据你的需求来处理
	}
}

// 批量生成输出，参数示例：/app/output/image/2025-05-01/1746093311_4_%04d.png
func BatchGenerateDownloadURL(fileFormat string) []string {
	images := make([]string, 0)

	for i := 1; i < 10000; i++ {
		file := fmt.Sprintf(fileFormat, i)
		if _, err := os.Stat(file); err != nil {
			break
		}

		newFile := strings.Replace(file, consts.WebRootPrefix, config.URL, 1)
		if strings.HasPrefix(newFile, config.URL) {
			images = append(images, newFile)
		} else {
			panic(fmt.Sprintf("invalid local path, file: %s", file))
		}
	}

	return images
}
