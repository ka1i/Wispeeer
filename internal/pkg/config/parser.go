package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type config struct {
	filename string
	Error    error
}

func (c *config) Init() {
	options, err := configParser(c.filename)
	if err != nil {
		c.Error = err
		return
	}
	Cfg.Set(&options)
}

func (c *config) SetFile(filename string) {
	c.filename = filename
}

func (c *config) GetFile() string {
	return c.filename
}

var Configure = getConfig()

func getConfig() *config {
	return &config{
		filename: "config.yml",
		Error:    nil,
	}
}

func configParser(filename string) (Options, error) {
	options := Options{}

	configFile, err := os.ReadFile(filename)
	if err != nil {
		return options, err
	}

	err = yaml.Unmarshal(configFile, &options)
	if err != nil {
		return options, err
	}

	return options, err
}
