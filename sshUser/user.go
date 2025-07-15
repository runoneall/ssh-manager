package sshUser

func (u *SSHUsers) AddUser(
	name string,
	password string,
	isAdmin bool,
	servers []string,
	isDisable bool,
) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.items[name] = SSHUserItem{
		Name:      name,
		Password:  password,
		IsAdmin:   isAdmin,
		Servers:   servers,
		IsDisable: isDisable,
	}
}

func (u *SSHUsers) GetUser(name string) (SSHUserItem, bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	user, ok := u.items[name]
	return user, ok
}
