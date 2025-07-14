package sshShell

import (
	"fmt"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func StartShell(session ssh.Session) {

	// 创建终端
	terminal := term.NewTerminal(session, fmt.Sprintf("%s> ", session.User()))

	// 创建logger
	logger := func(msg string) {
		terminal.Write([]byte(msg + "\n"))
	}

	// 主循环
	for {
		line, err := terminal.ReadLine()

		// 无法读取命令行
		if err != nil {
			logger("读取命令行失败")
			session.Exit(1)
			return
		}

		// 退出
		if line == "exit" {
			session.Close()
			return
		}

		// 执行命令
		terminal.Write([]byte("Echo: " + line + "\n"))
	}

}
