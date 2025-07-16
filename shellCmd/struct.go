package shellCmd

import (
	"fmt"
	"ssh-manager/sshUser"
	"sync"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

type CommandItem struct {
	NeedAdmin bool
	Name      string
	Call      func(s ssh.Session, t *term.Terminal, arg []string)
	Help      string
}

type Commands struct {
	mu              sync.RWMutex
	cmds            []CommandItem
	defaultHandler  func(s ssh.Session, t *term.Terminal, arg []string)
	nonAdminHandler func(s ssh.Session, t *term.Terminal, arg []string)
}

// 全局命令管理器
var (
	globalCommandManager *Commands
	umanager             *sshUser.SSHUsers
	once                 sync.Once
)

func GetCommandManager() *Commands {
	once.Do(func() {
		umanager = sshUser.GetSSHUserManager()
		globalCommandManager = &Commands{
			cmds: make([]CommandItem, 0),
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
	})
	return globalCommandManager
}
