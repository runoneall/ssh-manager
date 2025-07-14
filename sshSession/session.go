package sshSession

import "github.com/gliderlabs/ssh"

func GetOnlineSessions() *OnlineSessions {
	return &currentOnlineSessions
}

func (s *OnlineSessions) AddSession(user string, session ssh.Session) {
	(*s)[user] = append((*s)[user], OnlineSessionItem{
		User:    user,
		Session: session,
	})
}

func (s *OnlineSessions) RemoveSession(user string, session ssh.Session) {
	for i, item := range (*s)[user] {
		if item.Session == session {
			(*s)[user] = append((*s)[user][:i], (*s)[user][i+1:]...)
			break
		}
	}
}

func (s *OnlineSessions) GetUserSessions(user string) []OnlineSessionItem {
	return (*s)[user]
}
