package sshUser

import (
	"fmt"
	"ssh-manager/helper"
)

func (users *SSHUsers) LoadFromJson(path string) error {
	if ok := helper.LoadJSON(path, users); !ok {
		return fmt.Errorf("无法从JSON文件 %s 中加载用户信息", path)
	}
	return nil
}

func (users *SSHUsers) SaveToJson(path string) error {
	helper.SaveJSON(path, users)
	return nil
}
