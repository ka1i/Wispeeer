package main

import (
	"log"
	"os"

	"github.com/wispeeer/wispeeer/internal/app"
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
		log.Println(err)
	}
}
