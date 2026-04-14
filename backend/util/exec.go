package util

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"

	"github.com/go-cmd/cmd"
)

func Exec(command string, args ...string) (string, string, error) {
	tmpCmd := exec.Command(command, args...)

	// 用于捕获命令的输出
	var out bytes.Buffer
	var errOut bytes.Buffer

	// 设置命令的输出
	tmpCmd.Stdout = &out
	tmpCmd.Stderr = &errOut

	// 执行命令
	err := tmpCmd.Run()
	if err != nil {
		return out.String(), errOut.String(), err
	}

	return out.String(), errOut.String(), nil
}

// SafeExec 安全执行外部命令，支持上下文超时控制和进程树终止
// 返回值: stdout, stderr, error
func SafeExec(ctx context.Context, command string, args ...string) (string, string, error) {
	// 创建输出缓冲区
	stdoutBuffer := bytes.NewBuffer(nil)
	stderrBuffer := bytes.NewBuffer(nil)

	// 创建命令选项
	options := cmd.Options{
		Buffered:  false, // 使用流式输出
		Streaming: true,  // 实时获取输出
	}

	// 创建命令对象
	c := cmd.NewCmdOptions(options, command, args...)

	// 启动命令（非阻塞）
	statusChan := c.Start()

	// 创建安全退出通道
	doneChan := make(chan struct{})
	defer close(doneChan)

	// 启动协程收集输出
	go func() {
		for {
			select {
			case line, open := <-c.Stdout:
				if !open {
					return
				}
				stdoutBuffer.WriteString(line + "\n")
			case line, open := <-c.Stderr:
				if !open {
					return
				}
				stderrBuffer.WriteString(line + "\n")
			case <-doneChan:
				return // 主函数退出时安全退出
			}
		}
	}()

	// 等待命令完成或上下文取消
	select {
	case finalStatus := <-statusChan:
		// 命令正常完成
		stdout := stdoutBuffer.String()
		stderr := stderrBuffer.String()

		if finalStatus.Error != nil {
			return stdout, stderr, fmt.Errorf("command failed: %w", finalStatus.Error)
		}
		return stdout, stderr, nil

	case <-ctx.Done():
		// 上下文取消（超时或手动取消）
		stdout := stdoutBuffer.String()
		stderr := stderrBuffer.String()

		c.Stop()

		return stdout, stderr, ctx.Err()
	}
}
