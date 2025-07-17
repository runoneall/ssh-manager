package sshShell

import (
	"ssh-manager/shellBin"
	"ssh-manager/shellConn"
	"ssh-manager/sshSession"
	"ssh-manager/sshUser"
)

var umanager *sshUser.SSHUsers = sshUser.GetSSHUserManager()
var smanager *sshSession.OnlineSessions = sshSession.GetSessionManager()
var bmanager *shellBin.Bins = shellBin.GetBinManager()
var cmanager *shellConn.Connects = shellConn.GetConnectManager()

func init() {

	bmanager.AddBin(false, "logout", binLogout, "更好的退出登录")
	bmanager.AddBin(false, "lg", binLogout, "logout的别名")
	bmanager.AddBin(false, "exit", binLogout, "logout的别名")

	bmanager.AddBin(true, "user-add", sbinUserAdd, "添加用户")
	bmanager.AddBin(true, "ua", sbinUserAdd, "user-add的别名")

	bmanager.AddBin(true, "user", sbinUserManage, "管理用户")
	bmanager.AddBin(true, "um", sbinUserManage, "user的别名")

	bmanager.AddBin(true, "user-list", sbinUserList, "列出所有用户")
	bmanager.AddBin(true, "ul", sbinUserList, "user-list的别名")

	bmanager.AddBin(false, "connect", binConnect, "连接到SHELL")
	bmanager.AddBin(false, "cd", binConnect, "connect的别名")

}
