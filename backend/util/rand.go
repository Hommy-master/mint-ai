package util

import (
	"math/rand"
	"sync"
	"time"
)

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ" // 36个字符

var (
	seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	mu         sync.Mutex
)

// RandString 生成指定长度的随机字符串（0-9, A-Z）
func RandString(length int) string {
	mu.Lock()
	defer mu.Unlock()

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}
