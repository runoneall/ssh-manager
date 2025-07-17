package shellBin

import (
	"fmt"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func (b *Bins) ShowHelp(s ssh.Session, t *term.Terminal) {
	b.mu.RLock()
	supported_bins := b.bins
	b.mu.RUnlock()

	isAdmin := umanager.IsAdmin(s.User())

	fmt.Fprintln(t, "可用命令:")
	fmt.Fprintln(t, "  help - 再次显示此帮助信息")

	for _, item := range supported_bins {
		// 只显示用户有权限查看的命令
		if !item.NeedAdmin || isAdmin {
			fmt.Fprintf(t,
				"  %s - %s\n",
				item.Name, item.Help,
			)
		}
	}
}
