package sshSession

import (
	"context"

	"github.com/gliderlabs/ssh"
	"github.com/google/uuid"
)

// 添加会话并返回自动清理函数
func (s *OnlineSessions) AddSession(user string, session ssh.Session) (string, context.CancelFunc) {
	s.mu.RLock()
	store, exists := s.users[user]
	s.mu.RUnlock()

	// 用户不存在时创建新存储
	if !exists {
		s.mu.Lock()
		// 双重检查避免竞态
		if store, exists = s.users[user]; !exists {
			store = &UserSessionStore{
				sessions: make(map[string]*OnlineSessionItem),
			}
			s.users[user] = store
		}
		s.mu.Unlock()
	}

	id := uuid.New().String()
	ctx, cancel := context.WithCancel(context.Background())

	store.mu.Lock()
	store.sessions[id] = &OnlineSessionItem{
		ID:      id,
		User:    user,
		Session: session,
		Ctx:     ctx,
	}
	store.mu.Unlock()

	// 返回会话ID和清理函数
	return id, func() {
		cancel() // 取消上下文
		s.RemoveSession(user, id)
	}
}

// 移除会话
func (s *OnlineSessions) RemoveSession(user string, id string) {
	s.mu.RLock()
	store, exists := s.users[user]
	s.mu.RUnlock()

	if !exists {
		return
	}

	store.mu.Lock()
	delete(store.sessions, id)

	// 清理空用户存储
	if len(store.sessions) == 0 {
		s.mu.Lock()
		delete(s.users, user)
		s.mu.Unlock()
	}
	store.mu.Unlock()
}

// 获取用户的所有会话
func (s *OnlineSessions) GetUserSessions(user string) map[string]*OnlineSessionItem {
	s.mu.RLock()
	store, exists := s.users[user]
	s.mu.RUnlock()

	if !exists {
		return nil
	}

	store.mu.RLock()
	defer store.mu.RUnlock()

	return store.sessions
}
