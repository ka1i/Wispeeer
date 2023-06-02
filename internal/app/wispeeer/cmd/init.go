package cmd

import (
	"fmt"

	"github.com/wispeeer/wispeeer/internal/pkg/tools"
	"github.com/wispeeer/wispeeer/pkg/assets"
	"github.com/wispeeer/wispeeer/pkg/logger"
	"github.com/wispeeer/wispeeer/pkg/utils"
)

func (c *CMD) Initialzation(title string) error {
	var err error

	if utils.IsExist(title) {
		return fmt.Errorf("%s: File exists", title)
	}

	logger.Task("init").Infof("wispeeer init %s", title)

	logger.Task("init").Info("unpkg embed assets")

	var storage = assets.GetStorage()
	fs := storage.Fs
	root := storage.Root
	err = tools.EmbedUnpkg(&fs, root, root, title)
	return err
}
