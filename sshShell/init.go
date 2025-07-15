package sshShell

import (
	"ssh-manager/shellCmd"
	"ssh-manager/sshSession"
)

var smanager *sshSession.OnlineSessions = sshSession.GetSessionManager()
var cmanager *shellCmd.Commands = shellCmd.GetCommandManager()

func init() {
	cmanager.AddCommand("exit", binExit, "退出登录")
	cmanager.AddCommand("logout", binLogout, "更好的退出登录")
	cmanager.AddCommand("lg", binLogout, "logout的别名")
}
