package sshShell

import (
	"fmt"
	"strings"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func StartShell(session ssh.Session) {

	// 创建终端
	ptyReq, winCh, isPty := session.Pty()
	if !isPty {
		fmt.Fprintln(session, "无法创建终端")
		session.Exit(1)
		return
	}
	terminal := term.NewTerminal(session, fmt.Sprintf("%s> ", session.User()))

	// 设置终端大小
	terminal.SetSize(ptyReq.Window.Width, ptyReq.Window.Height)
	go func() {
		for win := range winCh {
			terminal.SetSize(win.Width, win.Height)
		}
	}()

	// 主循环
	for {
		line, err := terminal.ReadLine()

		// 无法读取命令行
		if err != nil {
			fmt.Fprintln(terminal, "读取命令行失败")
			session.Exit(1)
			return
		}

		// 如果输入为空
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// 解析命令
		args := strings.Fields(line)
		cmd := args[0]

		// 显示帮助
		if cmd == "help" {
			cmanager.ShowHelp(terminal)
			continue
		}

		// 运行命令
		cmanager.RunCommand(cmd)(
			session, terminal, args,
		)

	}

}
