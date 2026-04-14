package util

import (
	"os"
	"testing"
)

// TestAppendToFile 测试 AppendToFile 函数
func TestAppendToFile(t *testing.T) {
	defer func() {
		_ = os.Remove("test.txt")
	}()

	err := AppendToFile("test.txt", "1.Hello, World!\n")
	if err != nil {
		t.Errorf("AppendToFile failed: %v", err)
	}
	err = AppendToFile("test.txt", "2.Hello, World!\n")
	if err != nil {
		t.Errorf("AppendToFile failed: %v", err)
	}
	err = AppendToFile("test.txt", "3.Hello, World!")
	if err != nil {
		t.Errorf("AppendToFile failed: %v", err)
	}
	err = AppendToFile("test.txt", "4.Hello, World!\n")
	if err != nil {
		t.Errorf("AppendToFile failed: %v", err)
	}
}
