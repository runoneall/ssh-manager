package sshSession

import (
	"context"
	"sync"

	"github.com/gliderlabs/ssh"
)

// 会话处理函数类型
type SessionHandler func(s ssh.Session)

// 自动管理会话生命周期的包装器
func (s *OnlineSessions) AutoSessionHandler(handler SessionHandler) ssh.Handler {
	return func(session ssh.Session) {
		user := session.User()

		// 添加会话并获取清理函数
		_, cleanup := s.AddSession(user, session)

		// 确保清理函数只执行一次
		var once sync.Once
		cleanupWrapper := func() {
			once.Do(cleanup)
		}

		// 双重保障清理机制
		defer cleanupWrapper() // 正常退出时清理

		// 监听会话结束信号
		go func() {
			<-session.Context().Done() // 等待会话结束
			cleanupWrapper()           // 异常退出时清理
		}()

		// 执行实际的会话处理逻辑
		handler(session)
	}
}

// 自定义会话类型支持上下文
type sessionWithContext struct {
	ssh.Session
	ctx context.Context
}

func (s sessionWithContext) Context() context.Context {
	return s.ctx
}
