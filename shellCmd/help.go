package shellCmd

import (
	"fmt"

	"golang.org/x/term"
)

func (c *Commands) ShowHelp(t *term.Terminal) {
	c.mu.RLock()
	supported_cmds := c.cmds
	c.mu.RUnlock()

	fmt.Fprintln(t, "可用命令:")
	fmt.Fprintln(t, "  help - 再次显示此帮助信息")
	for _, item := range supported_cmds {
		fmt.Fprintf(t,
			"  %s - %s\n",
			item.Name, item.Help,
		)
	}
}
