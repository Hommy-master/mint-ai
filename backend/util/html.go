package util

import (
	"strings"

	"golang.org/x/net/html"
)

func IsHTML(content string) bool {
	// 解析HTML内容
	doc, err := html.Parse(strings.NewReader(content))
	return err == nil && doc != nil
}
