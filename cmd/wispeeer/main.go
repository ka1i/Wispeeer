package main

import (
	"os"

	"github.com/wispeeer/wispeeer/internal/app"
	"github.com/wispeeer/wispeeer/pkg/logger"
)

func main() {
	// bootstrap
	service := app.App
	service.Usage = app.Wispeeer.Usage()
	service.Flags = app.Wispeeer.Flags()
	service.Service = app.Wispeeer.Service()
	code, err := service.Run()
	defer os.Exit(code)
	if err != nil {
		logger.Task("app").Error(err)
	}
}
