package sshShell

import (
	"ssh-manager/shellCmd"
	"ssh-manager/sshSession"
	"ssh-manager/sshUser"
)

var umanager *sshUser.SSHUsers = sshUser.GetSSHUserManager()
var smanager *sshSession.OnlineSessions = sshSession.GetSessionManager()
var cmanager *shellCmd.Commands = shellCmd.GetCommandManager()

func init() {

	cmanager.AddCommand("logout", binLogout, "更好的退出登录")
	cmanager.AddCommand("lg", binLogout, "logout的别名")
	cmanager.AddCommand("exit", binLogout, "logout的别名")

	cmanager.AddCommand("user-add", sbinUserAdd, "添加用户")
	cmanager.AddCommand("ua", sbinUserAdd, "user-add的别名")

	cmanager.AddCommand("user", sbinUserManage, "管理用户")
	cmanager.AddCommand("um", sbinUserManage, "user的别名")

	cmanager.AddCommand("user-list", sbinUserList, "列出所有用户")
	cmanager.AddCommand("ul", sbinUserList, "user-list的别名")

}
