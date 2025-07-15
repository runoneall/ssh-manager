package sshSession

import (
	"context"
	"sync"

	"github.com/gliderlabs/ssh"
)

// 用户会话存储
type UserSessionStore struct {
	mu       sync.RWMutex
	sessions map[string]*OnlineSessionItem // sessionID -> session
}

// 在线会话项
type OnlineSessionItem struct {
	ID      string
	User    string
	Session ssh.Session
	Ctx     context.Context // 会话上下文
}

// 全局会话管理器
type OnlineSessions struct {
	mu    sync.RWMutex
	users map[string]*UserSessionStore // username -> store
}

// 全局会话管理器实例
var globalSessionManager = &OnlineSessions{
	users: make(map[string]*UserSessionStore),
}

// 获取会话管理器
func GetSessionManager() *OnlineSessions {
	return globalSessionManager
}
