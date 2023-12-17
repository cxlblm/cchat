package provider

import (
	"cchart/internal/kernel"
	"github.com/BurntSushi/toml"
)

const defaultConfigPath = "config.toml"

func path() string {
	return defaultConfigPath
}

func NewConfig() *kernel.Config {
	c := new(kernel.Config)
	_, err := toml.DecodeFile(path(), c)

	if err != nil {
		return nil
	}

	return c
}
