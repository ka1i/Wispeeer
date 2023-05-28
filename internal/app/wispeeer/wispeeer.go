package wispeeer

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/wispeeer/wispeeer/internal/app/wispeeer/cmd"
	"github.com/wispeeer/wispeeer/pkg/logger"
	"github.com/wispeeer/wispeeer/pkg/utils"
)

func Usage() string {
	var buffer strings.Builder

	buffer.WriteString("  -i, init          Create a new blog. (e.g wispeeer -i <alias>)")

	buffer.WriteString("\n")
	return buffer.String()
}

func Flags() (bool, error) {
	var err error
	var argv = os.Args[1:]
	var argc = len(argv)
	var isBreak bool = false

	run := cmd.Run()

	switch argv[0] {
	case "-i", "init":
		isBreak = true
		if argc >= 2 {
			if utils.IsValid(argv[1]) {
				err = run.Initialzation(argv[1])
			} else {
				err = fmt.Errorf("invalid blog folder name")
			}
		} else {
			err = fmt.Errorf("wispeeer init <ka1i.github.io>")
		}
	default:
		err = errors.New("usage: wispeeer -h")
	}
	return isBreak, err
}

func Service() error {
	var err error

	logger.Task("app").Info("usage: wispeeer -h")

	return err
}
