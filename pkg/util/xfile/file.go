package xfile

import (
	"os"
	"runtime"
	"strings"
)

// IsDirectory ...
func IsDirectory(path string) (bool, error) {
	f, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	switch mode := f.Mode(); {
	case mode.IsDir():
		return true, nil
	case mode.IsRegular():
		return false, nil
	}
	return false, nil
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	if runtime.GOOS == "windows" {
		dirctory = strings.Replace(dirctory, "\\", "/", -1)
	}
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

// CheckAndGetParentDir ...
func CheckAndGetParentDir(path string) string {
	// check path is the directory
	isDir, err := IsDirectory(path)
	if err != nil || isDir {
		return path
	}
	return getParentDirectory(path)
}
