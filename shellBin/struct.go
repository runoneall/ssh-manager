package shellBin

import (
	"fmt"
	"ssh-manager/sshUser"
	"sync"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

type BinItem struct {
	NeedAdmin bool
	Name      string
	Call      func(s ssh.Session, t *term.Terminal, arg []string)
	Help      string
}

type Bins struct {
	mu              sync.RWMutex
	bins            []BinItem
	defaultHandler  func(s ssh.Session, t *term.Terminal, arg []string)
	nonAdminHandler func(s ssh.Session, t *term.Terminal, arg []string)
}

var globalBinManager *Bins = &Bins{
	bins: make([]BinItem, 0),
	defaultHandler: func(s ssh.Session, t *term.Terminal, arg []string) {
		if len(arg) > 0 {
			fmt.Fprintln(t, "未找到命令:", arg[0])
		}
		fmt.Fprintln(t, "请使用 help 命令查看可用命令")
	},
	nonAdminHandler: func(s ssh.Session, t *term.Terminal, arg []string) {
		fmt.Fprintln(t, "你不能这么做!")
	},
}

var umanager *sshUser.SSHUsers = sshUser.GetSSHUserManager()

func GetBinManager() *Bins {
	return globalBinManager
}
