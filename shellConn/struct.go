package shellConn

import (
	"fmt"
	"sync"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

type ConnInfo struct {
	Type              string `json:"type"` // local or ssh
	Name              string `json:"name"`
	Local_Shell       string `json:"local_shell"`
	SSH_Host          string `json:"ssh_host"`
	SSH_Port          string `json:"ssh_port"`
	SSH_User          string `json:"ssh_user"`
	SSH_Type          string `json:"ssh_type"` // password or key
	SSH_Type_Password string `json:"ssh_type_password"`
	SSH_Type_Key      string `json:"ssh_type_key"`
}

type shellCall func(i ConnInfo, s ssh.Session, t *term.Terminal)

type Connects struct {
	mu               sync.RWMutex
	items            map[string]ConnInfo
	localShellCall   shellCall
	sshShellCall     shellCall
	unknownShellCall shellCall
}

var globalConnectManager = &Connects{
	items: make(map[string]ConnInfo),
	localShellCall: func(i ConnInfo, s ssh.Session, t *term.Terminal) {
		fmt.Fprintln(t, "暂不支持本地连接")
	},
	sshShellCall: func(i ConnInfo, s ssh.Session, t *term.Terminal) {
		fmt.Fprintln(t, "暂不支持SSH连接")
	},
	unknownShellCall: func(i ConnInfo, s ssh.Session, t *term.Terminal) {
		fmt.Fprintln(t, "未知连接")
	},
}

func GetConnectManager() *Connects {
	return globalConnectManager
}
