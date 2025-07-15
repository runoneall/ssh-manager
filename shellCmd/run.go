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

func (c *Commands) RunCommand(
	cmd string,
) func(s ssh.Session, t *term.Terminal, arg []string) {
	c.mu.RLock()
	supported_cmds := c.cmds
	c.mu.RUnlock()

	for _, item := range supported_cmds {
		if item.Name == cmd {
			return item.Call
		}
	}
	return c.defaultHandler
}
