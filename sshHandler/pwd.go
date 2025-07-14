package sshHandler

import (
	"ssh-manager/sshUser"

	"github.com/gliderlabs/ssh"
)

func OnPasswordAuth(ctx ssh.Context, password string) bool {
	user := ctx.User()
	userConfig := sshUser.GetSSHUsers().GetUser(user)

	// 如果 用户不存在 || 用户被禁用 || 密码错误
	if userConfig == nil || userConfig.IsDisable || password != userConfig.Password {
		return false
	}

	return true
}
