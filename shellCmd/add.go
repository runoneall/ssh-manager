package shellCmd

import (
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func (c *Commands) AddCommand(
	needAdmin bool,
	name string,
	call func(s ssh.Session, t *term.Terminal, arg []string),
	help string,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cmds = append(c.cmds, CommandItem{
		NeedAdmin: needAdmin,
		Name:      name,
		Call:      call,
		Help:      help,
	})
}
