package sshSession

import "ssh-manager/helper"

// 统计用户的会话数量
func (s *OnlineSessions) CountSession(user string) int {
	s.mu.RLock()
	store, exists := s.users[user]
	s.mu.RUnlock()

	if !exists {
		return 0
	}

	store.mu.RLock()
	defer store.mu.RUnlock()
	return len(store.sessions)
}

// 统计用户的唯一IP数量
func (s *OnlineSessions) CountIP(user string) int {
	s.mu.RLock()
	store, exists := s.users[user]
	s.mu.RUnlock()

	if !exists {
		return 0
	}

	ipSet := make(map[string]struct{})
	store.mu.RLock()
	defer store.mu.RUnlock()

	for _, item := range store.sessions {
		ip := helper.GetClientIP(item.Session.RemoteAddr())
		ipSet[ip] = struct{}{}
	}
	return len(ipSet)
}
