package sshSession

import (
	"net"
	"slices"

	"github.com/gliderlabs/ssh"
)

func getClientIP(s ssh.Session) string {
	addr := s.RemoteAddr().String()
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return addr
	}
	return host
}

func (s *OnlineSessions) CountSession(user string) int {
	return len(s.GetUserSessions(user))
}

func (s *OnlineSessions) CountIP(user string) int {
	ipList := []string{}
	for _, item := range (*s)[user] {
		ip := getClientIP(item.Session)
		if !slices.Contains(ipList, ip) {
			ipList = append(ipList, ip)
		}
	}
	return len(ipList)
}
