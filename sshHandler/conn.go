package sshHandler

import (
	"fmt"
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

	// 欢迎消息
	fmt.Fprintln(s, "* Welcome to SSH Manager!")
	fmt.Fprintf(
		s, "* 已作为 %s 用户登录, 已有 %d 个会话在 %d 个IP上\n",
		u, manager.CountSession(u), manager.CountIP(u),
	)

	// 进入shell
	sshShell.StartShell(s)

}
