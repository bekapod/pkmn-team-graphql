package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppHost      string
	Port         string
	DatabaseUrl  string
	DatabaseName string
}

func New() Config {
	return Config{
		AppHost:      os.Getenv("APP_HOST"),
		Port:         os.Getenv("PORT"),
		DatabaseUrl:  os.Getenv("DATABASE_URL"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
	}
}

func (c Config) AppHostWithPort() string {
	return fmt.Sprintf("%s:%s", c.AppHost, c.Port)
}
