package sshSession

import "github.com/gliderlabs/ssh"

type OnlineSessionItem struct {
	User    string
	Session ssh.Session
}

type OnlineSessions map[string][]OnlineSessionItem

var currentOnlineSessions OnlineSessions = make(OnlineSessions)
