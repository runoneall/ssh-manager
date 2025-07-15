package sshShell

import (
	"fmt"

	"github.com/akamensky/argparse"
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func binExit(s ssh.Session, t *term.Terminal, arg []string) {
	s.Exit(0)
}

func binLogout(s ssh.Session, t *term.Terminal, arg []string) {

	// 直接退出
	if TokenAt(arg, 1) == "" {
		s.Exit(0)
		return
	}

	// 创建解析器
	parser := argparse.NewParser("logout", "更好的退出登录")
	parser.DisableHelp()

	// 添加help选项
	isHelp := parser.Flag("h", "help", &argparse.Options{Help: "显示帮助信息"})

	// 添加all选项
	isAll := parser.Flag("a", "all", &argparse.Options{Help: "退出所有会话"})

	// 未知参数
	err := parser.Parse(arg)
	if err != nil {
		fmt.Fprint(t, parser.Usage(err))
		return
	}

	switch {

	// 显示帮助信息
	case *isHelp:
		fmt.Fprint(t, parser.Usage(nil))

	// 退出所有会话
	case *isAll:
		for _, store := range smanager.GetUserSessions(s.User()) {
			store.Session.Exit(0)
		}

	}
}
