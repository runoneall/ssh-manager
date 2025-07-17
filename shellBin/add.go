package shellBin

import (
	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func (b *Bins) AddBin(
	needAdmin bool,
	name string,
	call func(s ssh.Session, t *term.Terminal, arg []string),
	help string,
) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.bins = append(b.bins, BinItem{
		NeedAdmin: needAdmin,
		Name:      name,
		Call:      call,
		Help:      help,
	})
}
