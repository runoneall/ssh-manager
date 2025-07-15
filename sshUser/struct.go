package sshUser

import "sync"

type SSHUserItem struct {
	Name      string   `json:"name"`
	Password  string   `json:"password"`
	IsAdmin   bool     `json:"isAdmin"`
	Servers   []string `json:"servers"`
	IsDisable bool     `json:"isDisable"`
}

type SSHUsers struct {
	mu    sync.RWMutex
	items map[string]SSHUserItem
}

var globalSSHUsers = &SSHUsers{
	items: make(map[string]SSHUserItem),
}

func GetSSHUserManager() *SSHUsers {
	return globalSSHUsers
}
