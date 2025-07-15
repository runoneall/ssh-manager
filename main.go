package main

import (
	"fmt"
	"ssh-manager/helper"
	"ssh-manager/sshHandler"
	"ssh-manager/sshServer"
	"ssh-manager/sshUser"
	"ssh-manager/vars"
)

var umanager *sshUser.SSHUsers = sshUser.GetSSHUserManager()

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

	// 创建用户配置
	userConfigPath := vars.FILE_USER_CONFIG
	if !helper.IsExist(userConfigPath) {
		fmt.Println("创建用户配置文件...")

		// 添加超级用户
		umanager.AddUser("admin", "admin", true, []string{}, false)
		if !umanager.SaveToJson(userConfigPath) {
			panic(fmt.Errorf("不能创建用户配置文件"))
		}
		fmt.Println("默认超级用户: admin, 密码: admin (该信息只会显示一次!)")
	}

	// 加载用户配置
	fmt.Println("加载用户信息...")
	umanager.LoadFromJson(userConfigPath)

	// 启动SSH服务器
	sshServer.Start(
		sshHandler.OnPasswordAuth,
		sshHandler.OnConnect,
	)

}
