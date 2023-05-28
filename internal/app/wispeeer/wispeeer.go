package wispeeer

import (
	"errors"
	"log"
	"os"
	"strings"
)

func Usage() string {
	var buffer strings.Builder

	return buffer.String()
}

func Flags() (bool, error) {
	var err error
	var argv = os.Args[1:]
	var isBreak bool = false

	switch argv[0] {
	default:
		err = errors.New("please check usage")
	}
	return isBreak, err
}

func Service() error {
	var err error

	log.Println("Sorry not allow exec")

	return err
}
