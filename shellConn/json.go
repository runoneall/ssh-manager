package shellConn

import (
	"fmt"
	"ssh-manager/helper"
)

func (c *Connects) LoadFromJson(path string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	tempItems := make(map[string]ConnInfo)
	if ok := helper.LoadJSON(path, &tempItems); !ok {
		return fmt.Errorf("无法从JSON文件 %s 中加载用户信息", path)
	}
	c.items = tempItems
	return nil
}

func (c *Connects) SaveToJson(path string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return helper.SaveJSON(path, c.items)
}
