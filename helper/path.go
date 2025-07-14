package helper

import "os"

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func DeleteFileIfExist(path string) bool {
	if !IsExist(path) {
		return true
	}
	if err := os.Remove(path); err != nil {
		return false
	}
	return true
}
