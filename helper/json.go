package helper

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveJSON(path string, data any) bool {
	jsonData, _ := json.MarshalIndent(data, "", "  ")

	if err := os.WriteFile(path, jsonData, 0644); err != nil {
		fmt.Println("不能写入文件:", path)
		return false
	}
	return true
}

func LoadJSON(path string, target any) bool {
	jsonData, err := os.ReadFile(path)

	if err != nil {
		fmt.Println("不能读取文件:", path)
		return false
	}

	if err := json.Unmarshal(jsonData, target); err != nil {
		fmt.Println("不能解析JSON:", err)
		return false
	}

	return true
}
