package shellCmd

import (
	"fmt"
	"sync"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

type CommandItem struct {
	Name string
	Call func(s ssh.Session, t *term.Terminal, arg []string)
	Help string
}

type Commands struct {
	mu             sync.RWMutex
	cmds           []CommandItem
	defaultHandler func(s ssh.Session, t *term.Terminal, arg []string)
}

var globalCommandManager = &Commands{
	cmds: make([]CommandItem, 0),
	defaultHandler: func(s ssh.Session, t *term.Terminal, arg []string) {
		fmt.Fprintln(t, "* 未找到命令")
		fmt.Fprintln(t, "* 请使用 help 命令查看可用命令")
	},
}

func GetCommandManager() *Commands {
	return globalCommandManager
}
