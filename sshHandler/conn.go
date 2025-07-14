package sshHandler

import (
	"fmt"
	"ssh-manager/sshSession"
	"ssh-manager/sshShell"

	"github.com/gliderlabs/ssh"
)

func OnConnect(session ssh.Session) {
	allSessions := sshSession.GetOnlineSessions()
	allSessions.AddSession(session.User(), session)
	currentUser := session.User()

	// 控制台消息
	fmt.Printf("用户 %s(%s) 已连接\n", currentUser, session.RemoteAddr())

	// 连接关闭时清理
	defer func() {
		fmt.Printf("用户 %s(%s) 已断开连接\n", currentUser, session.RemoteAddr())
		allSessions.RemoveSession(currentUser, session)
	}()

	// 封装ssh连接logger
	sshLogger := func(msg string) {
		session.Write([]byte(msg + "\n"))
	}

	// 显示欢迎消息
	sshLogger(fmt.Sprintf(
		"SSH Manager: 已作为 %s 用户登录", currentUser,
	))
	sshLogger(fmt.Sprintf(
		"SSH Manager: 已有 %d 个会话在 %d 个IP上",
		allSessions.CountSession(currentUser),
		allSessions.CountIP(currentUser),
	))

	// 进入主循环
	sshShell.StartShell(session, sshLogger)
}
