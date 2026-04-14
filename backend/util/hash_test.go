package util

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	str := "Hello, World!"
	hash := MD5(str)
	fmt.Printf("MD5(%s) = %s\n", str, hash)
}
