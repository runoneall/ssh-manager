package sshHandler

import (
	"fmt"
	"ssh-manager/helper"
	"ssh-manager/sshUser"

	"github.com/gliderlabs/ssh"
)

var umanager *sshUser.SSHUsers = sshUser.GetSSHUserManager()

func OnPasswordAuth(ctx ssh.Context, password string) bool {
	user := ctx.User()
	userConfig, isExist := umanager.GetUser(user)

	// 如果 用户不存在 || 用户被禁用 || 密码错误
	if !isExist || userConfig.IsDisable || password != userConfig.Password {
		fmt.Printf(
			"用户 %s(%s) 使用了错误的密码 %s 尝试登录\n",
			user, helper.GetClientIP(ctx.RemoteAddr()), password,
		)
		return false
	}

	return true
}
