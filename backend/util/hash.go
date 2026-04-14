package util

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/gogf/gf/v2/os/glog"
)

// 计算字符串的MD5值
func MD5(str string) string {
	h := md5.New()

	// 写入数据到哈希对象
	if _, err := io.WriteString(h, str); err != nil {
		glog.Errorf(nil, "failed to write to hash object: %v", err)
		return ""
	}

	// 计算哈希值并转换为十六进制字符串
	return fmt.Sprintf("%x", h.Sum(nil))
}
