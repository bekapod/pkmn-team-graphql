package config

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

type Config struct {
	AppHost     string
	Port        string
	DatabaseUrl string
}

func New() Config {
	return Config{
		AppHost:     os.Getenv("APP_HOST"),
		Port:        os.Getenv("PORT"),
		DatabaseUrl: os.Getenv("DATABASE_URL"),
	}
}

func (c Config) AppHostWithPort() string {
	return fmt.Sprintf("%s:%s", c.AppHost, c.Port)
}

func (c Config) DatabaseName() string {
	u, err := url.Parse(c.DatabaseUrl)

	if err != nil {
		panic(err)
	}

	return strings.Replace(u.Path, "/", "", 1)
}
