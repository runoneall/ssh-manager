package shellCmd

import (
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func (c *Commands) SetDefaultHandler(
	handler func(s ssh.Session, t *term.Terminal, arg []string),
) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.defaultHandler = handler
}

func (c *Commands) SetNonAdminHandler(
	handler func(s ssh.Session, t *term.Terminal, arg []string),
) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.nonAdminHandler = handler
}

func (c *Commands) RunCommand(
	cmd string,
) func(s ssh.Session, t *term.Terminal, arg []string) {
	c.mu.RLock()
	supported_cmds := c.cmds
	c.mu.RUnlock()

	for _, item := range supported_cmds {
		if item.Name == cmd {
			return func(s ssh.Session, t *term.Terminal, arg []string) {
				// 检查管理员权限
				if item.NeedAdmin && !umanager.IsAdmin(s.User()) {
					c.nonAdminHandler(s, t, arg)
					return
				}
				item.Call(s, t, arg)
			}
		}
	}
	return c.defaultHandler
}
