package sshShell

import (
	"fmt"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func StartShell(session ssh.Session, logger func(msg string)) {

	// 获取PTY
	_, _, isPty := session.Pty()
	if !isPty {
		logger("无法获取PTY")
		session.Exit(1)
		return
	}

	// 创建终端
	terminal := term.NewTerminal(session, fmt.Sprintf("%s> ", session.User()))
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
