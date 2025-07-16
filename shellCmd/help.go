package shellCmd

import (
	"fmt"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func (c *Commands) ShowHelp(s ssh.Session, t *term.Terminal) {
	c.mu.RLock()
	supported_cmds := c.cmds
	c.mu.RUnlock()

	isAdmin := umanager.IsAdmin(s.User())

	fmt.Fprintln(t, "可用命令:")
	fmt.Fprintln(t, "  help - 再次显示此帮助信息")

	for _, item := range supported_cmds {
		// 只显示用户有权限查看的命令
		if !item.NeedAdmin || isAdmin {
			fmt.Fprintf(t,
				"  %s - %s\n",
				item.Name, item.Help,
			)
		}
	}
}
