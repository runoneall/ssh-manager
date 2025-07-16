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

func (u *SSHUsers) IsAdmin(name string) bool {
	user, ok := u.GetUser(name)
	if !ok {
		return false
	}
	return user.IsAdmin
}

func (u *SSHUsers) IsExist(name string) bool {
	_, ok := u.GetUser(name)
	return ok
}

func (u *SSHUsers) DeleteUser(name string) {
	u.mu.Lock()
	defer u.mu.Unlock()

	delete(u.items, name)
}

func (u *SSHUsers) ListUser() []string {
	u.mu.RLock()
	defer u.mu.RUnlock()

	var users []string
	for name := range u.items {
		users = append(users, name)
	}
	return users
}
