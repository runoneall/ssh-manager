package sshShell

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"os/user"
	"sync"
	"time"

	"github.com/creack/pty"
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func binConnect(s ssh.Session, t *term.Terminal, arg []string) {
	fmt.Fprintln(t, "正在连接...")

	// 创建终端
	ptyReq, winCh, isPty := s.Pty()
	if !isPty {
		fmt.Fprintln(t, "无法创建终端")
		return
	}

	// 创建上下文
	ctx, cancel := context.WithCancel(s.Context())
	defer cancel()

	// 启动子进程
	cmd := exec.Command("/bin/bash")
	u, err := user.Current()
	if err != nil {
		fmt.Println("获取当前用户失败:", err)
		fmt.Fprintln(s, "获取当前用户失败")
	}
	cmd.Dir = u.HomeDir

	// 启动终端
	ptyFile, err := pty.StartWithAttrs(cmd, &pty.Winsize{
		Rows: uint16(ptyReq.Window.Height),
		Cols: uint16(ptyReq.Window.Width),
	}, nil)
	if err != nil {
		fmt.Println("启动失败:", err)
		fmt.Fprintln(s, "启动失败")
		return
	}

	// 使用WaitGroup确保所有协程退出
	var wg sync.WaitGroup

	// 监听窗口变化
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case win, ok := <-winCh:
				if !ok {
					return // 通道关闭时退出
				}
				if err := pty.Setsize(ptyFile, &pty.Winsize{
					Rows: uint16(win.Height),
					Cols: uint16(win.Width),
				}); err != nil {
					fmt.Println("调整终端大小失败:", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	// 标准输入
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, _ = io.Copy(ptyFile, s)
		cancel() // 输入结束时取消上下文
	}()

	// 标准输出
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, _ = io.Copy(s, ptyFile)
		cancel() // 输出结束时取消上下文
	}()

	// 等待命令退出
	cmd.Wait()

	// 关闭伪终端以释放资源
	ptyFile.Close()

	// 设置超时防止永久阻塞
	waitDone := make(chan struct{})
	go func() {
		wg.Wait()
		close(waitDone)
	}()

	select {
	case <-waitDone:
		// 正常结束
	case <-time.After(1 * time.Second):
		// 超时后强制终止进程
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
	}

	// 重置终端大小
	t.SetSize(ptyReq.Window.Width, ptyReq.Window.Height)
}
