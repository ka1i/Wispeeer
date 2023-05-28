package cmd

import (
	"fmt"

	"github.com/wispeeer/wispeeer/pkg/logger"
	"github.com/wispeeer/wispeeer/pkg/utils"
)

func (c *CMD) Initialzation(title string) error {
	var err error

	if utils.IsExist(title) {
		return fmt.Errorf("%s: File exists", title)
	}

	logger.Task("init").Infof("wispeeer init %s", title)

	return err
}
