package initHelper

import (
	"errors"
	"os"
)

var FolderAlreadyExistErr error = errors.New("Folder already exists")

func CreateFolder(path string) error {
	_, err := os.Stat(path)

	// 不存在则新建
	if os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}

	// 存在则错误
	return FolderAlreadyExistErr
}
