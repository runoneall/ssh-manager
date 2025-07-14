package helper

import (
	"errors"
	"fmt"
	"os"
)

var FolderAlreadyExistErr error = errors.New("Folder already exists")

func CreateFolder(path string) error {

	// 不存在则新建
	if !IsExist(path) {
		fmt.Println("创建目录:", path)
		return os.MkdirAll(path, 0755)
	}

	// 存在则错误
	return FolderAlreadyExistErr
}
