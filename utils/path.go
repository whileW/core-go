package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//创建目录
func MkdirIfNotExist(dir string) error {
	if len(dir) == 0 {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return nil
}

// @title    createDir
// @description   获取运行地址
// @auth                     （2020/04/05  20:22）
// @param     path            string
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("create GetCurrentDirectory ", err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//获取app名称
func GetAppname() string {
	full := os.Args[0]
	full = strings.Replace(full, "\\", "/", -1)
	splits := strings.Split(full, "/")
	if len(splits) >= 1 {
		name := splits[len(splits)-1]
		name = strings.TrimSuffix(name, ".exe")
		return name
	}
	return ""
}