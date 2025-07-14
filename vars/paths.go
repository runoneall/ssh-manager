package vars

import (
	"os"
	"path"
)

func getProgramRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}

// 目录定义
var FOLDER_PROGRAM_ROOT string = getProgramRoot()
var FOLDER_CONFIG = path.Join(FOLDER_PROGRAM_ROOT, "config")
var FOLDER_SSH_SERVER_KEYS = path.Join(FOLDER_PROGRAM_ROOT, "keys")

// 文件定义
var FILE_SERVER_CONFIG = path.Join(FOLDER_CONFIG, "server.json")
var FILE_USER_CONFIG = path.Join(FOLDER_CONFIG, "user.json")
var FILE_SSH_CONNECTION_CONFIG = path.Join(FOLDER_CONFIG, "ssh_connections.json")
var FILE_SSH_SERVER_PRIVATE_KEY = path.Join(FOLDER_SSH_SERVER_KEYS, "id_rsa")
var FILE_SSH_SERVER_PUBLIC_KEY = path.Join(FOLDER_SSH_SERVER_KEYS, "id_rsa.pub")
