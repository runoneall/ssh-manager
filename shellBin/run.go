package shellBin

import (
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func (b *Bins) SetDefaultHandler(
	handler func(s ssh.Session, t *term.Terminal, arg []string),
) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.defaultHandler = handler
}

func (b *Bins) SetNonAdminHandler(
	handler func(s ssh.Session, t *term.Terminal, arg []string),
) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.nonAdminHandler = handler
}

func (b *Bins) RunBin(
	bin string,
) func(s ssh.Session, t *term.Terminal, arg []string) {
	b.mu.RLock()
	supported_bins := b.bins
	b.mu.RUnlock()

	for _, item := range supported_bins {
		if item.Name == bin {
			return func(s ssh.Session, t *term.Terminal, arg []string) {
				// 检查管理员权限
				if item.NeedAdmin && !umanager.IsAdmin(s.User()) {
					b.nonAdminHandler(s, t, arg)
					return
				}
				item.Call(s, t, arg)
			}
		}
	}
	return b.defaultHandler
}
