package utils

import (
	"os"
	"strings"

	"github.com/wispeeer/wispeeer/pkg/logger"
)

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func GetWorkspace() string {
	dir, err := os.Getwd()
	if err != nil {
		logger.Task("app").Error(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsFile(path string) bool {
	return !IsDir(path)
}
