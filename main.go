package main

import (
	"ssh-manager/initHelper"
)

func main() {

	// 初始化文件夹
	for _, folderPath := range []string{
		CONFIG_FOLDER,
	} {
		err := initHelper.CreateFolder(folderPath)
		if err != nil && err != initHelper.FolderAlreadyExistErr {
			panic(err)
		}
	}

}
