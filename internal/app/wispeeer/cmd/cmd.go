package cmd

import (
	"github.com/wispeeer/wispeeer/internal/pkg/config"
	"github.com/wispeeer/wispeeer/pkg/logger"
	"github.com/wispeeer/wispeeer/pkg/utils"
)

type CMD struct {
	Articles []string
	Options  config.Options
}

func Run() *CMD {
	logger.Task("wispeeer").Infof("workspace : %v", utils.GetWorkspace())
	return &CMD{
		Articles: make([]string, 0),
		Options:  *config.Cfg.Get(),
	}
}
