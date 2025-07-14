package main

import (
	"fmt"
	"ssh-manager/helper"
	"ssh-manager/sshHandler"
	"ssh-manager/sshServer"
	"ssh-manager/sshUser"
	"ssh-manager/vars"
)

func main() {

	// 初始化文件夹
	for _, folderPath := range []string{
		vars.FOLDER_CONFIG,
		vars.FOLDER_SSH_SERVER_KEYS,
	} {
		err := helper.CreateFolder(folderPath)
		if err != nil && err != helper.FolderAlreadyExistErr {
			panic(err)
		}
	}

	// 添加超级用户
	userConfigPath := vars.FILE_USER_CONFIG
	if !helper.IsExist(userConfigPath) {
		fmt.Println("创建用户配置文件...")
		if err := sshUser.GetSSHUsers().SaveToJson(userConfigPath); err != nil {
			panic(fmt.Errorf("不能创建用户配置文件: %v", err))
		}
		fmt.Println("默认超级用户: admin, 密码: admin (该信息只会显示一次!)")
	}

	// 启动SSH服务器
	sshServer.Start(
		sshHandler.OnPasswordAuth,
		sshHandler.OnConnect,
	)

}
