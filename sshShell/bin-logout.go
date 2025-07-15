package sshShell

import (
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func binExit(s ssh.Session, t *term.Terminal, arg []string) {
	s.Exit(0)
}

func binLogout(s ssh.Session, t *term.Terminal, arg []string) {

	// 是否登出所有会话
	if TokenAt(arg, 0) == "all" {
		for _, store := range smanager.GetUserSessions(s.User()) {
			store.Session.Exit(0)
		}
	} else {
		// 否则只登出当前会话
		s.Exit(0)
	}

}
