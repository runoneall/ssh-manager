package sshUser

func GetSSHUsers() *SSHUsers {
	return &currentSSHUsers
}

func (users *SSHUsers) AddUser(
	name string,
	password string,
	isAdmin bool,
	servers []string,
	isDisable bool,
) {
	(*users)[name] = SSHUserItem{
		Name:      name,
		Password:  password,
		IsAdmin:   isAdmin,
		Servers:   servers,
		IsDisable: isDisable,
	}
}

func (users *SSHUsers) GetUser(name string) *SSHUserItem {
	if user, ok := (*users)[name]; ok {
		return &user
	}
	return nil
}
