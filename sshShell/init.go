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

	cmanager.AddCommand(false, "logout", binLogout, "更好的退出登录")
	cmanager.AddCommand(false, "lg", binLogout, "logout的别名")
	cmanager.AddCommand(false, "exit", binLogout, "logout的别名")

	cmanager.AddCommand(true, "user-add", sbinUserAdd, "添加用户")
	cmanager.AddCommand(true, "ua", sbinUserAdd, "user-add的别名")

	cmanager.AddCommand(true, "user", sbinUserManage, "管理用户")
	cmanager.AddCommand(true, "um", sbinUserManage, "user的别名")

	cmanager.AddCommand(true, "user-list", sbinUserList, "列出所有用户")
	cmanager.AddCommand(true, "ul", sbinUserList, "user-list的别名")

	cmanager.AddCommand(false, "connect", binConnect, "连接到SHELL")
	cmanager.AddCommand(false, "cd", binConnect, "connect的别名")

}
