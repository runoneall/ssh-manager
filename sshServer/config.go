package sshServer

import (
	"fmt"
	"ssh-manager/helper"
	"ssh-manager/vars"
)

func GetConfig() *SSHServerConfig {
	jsonFilePath := vars.FILE_SERVER_CONFIG

	// 不存在则新建
	if !helper.IsExist(jsonFilePath) {
		fmt.Println("配置文件不存在, 新建配置文件")
		config := &SSHServerConfig{
			Port: vars.VAR_SSH_SERVER_PORT,
		}
		if helper.SaveJSON(jsonFilePath, config) {
			return config
		}
		panic(fmt.Errorf("不能保存SSH服务器配置文件到:%s", jsonFilePath))
	}

	// 存在则读取
	config := &SSHServerConfig{}
	if helper.LoadJSON(jsonFilePath, config) {
		return config
	}
	panic(fmt.Errorf("不能读取SSH服务器配置文件:%s", jsonFilePath))

}
