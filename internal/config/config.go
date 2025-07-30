package config

import (
	"os"
)

type Config struct {
	Port string
}

func Load() (*Config, error) {
	return &Config{
		Port: os.Getenv("FORGEBASE_PORT"),
	}, nil
}
