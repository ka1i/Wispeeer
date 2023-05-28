package wispeeer

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/wispeeer/wispeeer/internal/app/wispeeer/cmd"
	"github.com/wispeeer/wispeeer/internal/pkg/config"
	"github.com/wispeeer/wispeeer/pkg/logger"
	"github.com/wispeeer/wispeeer/pkg/utils"
)

func Usage() string {
	var buffer strings.Builder

	buffer.WriteString("  -i, init          Create a new blog. (e.g wispeeer -i <alias>)\n")
	buffer.WriteString("  -n, new           Create a new post or page. (e.g wispeeer -n [page] <title>)\n")
	buffer.WriteString("  -s, server        Start the server. (e.g wispeeer -s)\n")
	buffer.WriteString("  -d, deploy        Deploy your website. (e.g wispeeer -d)\n")
	buffer.WriteString("  -g, generate      Generate static files. (e.g wispeeer -g)\n")
	buffer.WriteString("  -a, analyse       Analyse the blog detail. (e.g wispeeer -a)\n")

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
	case "-n", "new":
		if argc >= 2 {
			if config.Configure.Error == nil {
				if argv[1] == "page" && argc > 3 {
					if utils.IsValid(argv[2]) {
						err = run.NewPage(argv[2])
					}
				} else {
					if utils.IsValid(argv[1]) {
						err = run.NewPost(argv[1])
					}
				}
			} else {
				err = config.Configure.Error
			}
		} else {
			err = fmt.Errorf("wispeeer new [post] <title>")
		}
	case "-g", "generate":
		if config.Configure.Error == nil {
			err = run.Generate()
		} else {
			err = config.Configure.Error
		}
	case "-s", "server":
		if config.Configure.Error == nil {
			err = run.Server()
		} else {
			err = config.Configure.Error
		}
	case "-d", "deploy":
		if config.Configure.Error == nil {
			err = run.Deploy()
		} else {
			err = config.Configure.Error
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
