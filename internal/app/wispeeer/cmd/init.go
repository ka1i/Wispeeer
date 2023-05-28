package cmd

import (
	"fmt"
	"log"

	"github.com/wispeeer/wispeeer/pkg/utils"
)

// Initialzation ...
func (c *CMD) Initialzation(title string) error {
	var err error

	if utils.IsExist(title) {
		return fmt.Errorf("%s: File exists", title)
	}

	log.Println("wispeeer initialzation")

	return err
}
