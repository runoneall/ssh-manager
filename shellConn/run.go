package shellConn

func (c *Connects) StartShell(name string) shellCall {
	info, exist := c.GetConn(name)

	c.mu.RLock()
	defer c.mu.RUnlock()

	if !exist {
		return c.unknownShellCall
	}

	switch info.Type {
	case "local":
		return c.localShellCall
	case "ssh":
		return c.sshShellCall
	default:
		return c.unknownShellCall
	}
}
