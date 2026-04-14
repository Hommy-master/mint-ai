package util

import (
	"fmt"
	"testing"
)

func TestWriteTextToSRT(t *testing.T) {
	text := `
	1
	00:00:00,000 --> 00:00:05,000
	欢迎使用字幕编辑工具1
	Hello World

	2
	00:00:05,000 --> 00:00:10,000
	欢迎使用字幕编辑工具2
	Hello World

	3
	00:00:10,000 --> 00:00:15,000
	欢迎使用字幕编辑工具3
	Hello World

	4
	00:00:15,000 --> 00:00:20,000
	欢迎使用字幕编辑工具4
	Hello World

	5
	00:00:20,000 --> 00:00:25,000
	欢迎使用字幕编辑工具5
	Hello World
	`

	lines := ParseLines(text)
	for _, line := range lines {
		fmt.Printf("%s\n", line)
	}
}
