package cmd

import (
	"github.com/wispeeer/wispeeer/pkg/logger"
	"github.com/wispeeer/wispeeer/pkg/utils"
)

type CMD struct{}

func Run() *CMD {
	logger.Task("wispeeer").Infof("workspace : %v", utils.GetWorkspace())
	return &CMD{}
}
