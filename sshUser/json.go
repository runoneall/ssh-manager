package sshUser

import (
	"fmt"
	"ssh-manager/helper"
)

func (u *SSHUsers) LoadFromJson(path string) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	tempItems := make(map[string]SSHUserItem)
	if ok := helper.LoadJSON(path, &tempItems); !ok {
		return fmt.Errorf("无法从JSON文件 %s 中加载用户信息", path)
	}
	u.items = tempItems
	return nil
}

func (u *SSHUsers) SaveToJson(path string) bool {
	u.mu.Lock()
	defer u.mu.Unlock()

	return helper.SaveJSON(path, u.items)
}
