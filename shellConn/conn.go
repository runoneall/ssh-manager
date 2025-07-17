package shellConn

func (c *Connects) AddConn(info ConnInfo) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[info.Name] = info
}

func (c *Connects) GetConn(name string) (ConnInfo, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	conn, ok := c.items[name]
	return conn, ok
}

func (c *Connects) SetLocalShellCall(call shellCall) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.localShellCall = call
}

func (c *Connects) SetSSHShellCall(call shellCall) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.sshShellCall = call
}

func (c *Connects) SetUnknownShellCall(call shellCall) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.unknownShellCall = call
}
