package sshHandler

import (
	"fmt"
	"ssh-manager/helper"
	"ssh-manager/sshSession"
	"ssh-manager/sshShell"

	"github.com/gliderlabs/ssh"
)

var manager *sshSession.OnlineSessions = sshSession.GetSessionManager()

func OnConnect(session ssh.Session) {
	manager.AutoSessionHandler(handleConnection)(session)
}

func handleConnection(s ssh.Session) {
	u := s.User()
	ip := helper.GetClientIP(s.RemoteAddr())

	// 连接提示
	fmt.Printf("* 用户 %s(%s) 已连接\n", u, ip)

	// 欢迎消息
	fmt.Fprintln(s, "* Welcome to SSH Manager!")
	fmt.Fprintf(
		s, "* 已作为 %s 用户登录, 已有 %d 个会话在 %d 个IP上\n",
		u, manager.CountSession(u), manager.CountIP(u),
	)

	// 进入shell
	sshShell.StartShell(s)

	// 退出提示
	fmt.Printf("* 用户 %s(%s) 已断开连接\n", u, ip)

}
