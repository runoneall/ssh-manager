package sshShell

import (
	"fmt"
	"strings"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func StartShell(session ssh.Session) {

	// 创建终端
	terminal := term.NewTerminal(session, fmt.Sprintf("%s> ", session.User()))

	// 主循环
	for {
		line, err := terminal.ReadLine()
		line = strings.TrimSpace(line)

		// 无法读取命令行
		if err != nil {
			fmt.Fprintln(terminal, "读取命令行失败")
			session.Exit(1)
			return
		}

		// 如果输入为空，则跳过
		if line == "" {
			continue
		}

		// 解析命令
		parts := strings.Fields(line)
		cmd := parts[0]
		var args []string
		if len(parts) > 1 {
			args = parts[1:]
		} else {
			args = make([]string, 0)
		}

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
