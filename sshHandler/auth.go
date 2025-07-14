package sshHandler

import (
	"fmt"
	"ssh-manager/sshUser"

	"github.com/gliderlabs/ssh"
)

func OnPasswordAuth(ctx ssh.Context, password string) bool {
	user := ctx.User()
	userConfig := sshUser.GetSSHUsers().GetUser(user)

	// 如果 用户不存在 || 用户被禁用 || 密码错误
	if userConfig == nil || userConfig.IsDisable || password != userConfig.Password {
		fmt.Printf("用户 %s(%s) 使用了错误的密码 %s 尝试登录\n", user, ctx.RemoteAddr(), password)
		return false
	}

	return true
}
