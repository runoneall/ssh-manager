package sshUser

type SSHUserItem struct {
	Name      string   `json:"name"`
	Password  string   `json:"password"`
	IsAdmin   bool     `json:"isAdmin"`
	Servers   []string `json:"servers"`
	IsDisable bool     `json:"isDisable"`
}

type SSHUsers map[string]SSHUserItem

var currentSSHUsers SSHUsers = SSHUsers{
	"admin": SSHUserItem{
		Name:      "admin",
		Password:  "admin",
		IsAdmin:   true,
		Servers:   []string{},
		IsDisable: false,
	},
}
