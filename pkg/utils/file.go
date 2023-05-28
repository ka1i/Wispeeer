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
