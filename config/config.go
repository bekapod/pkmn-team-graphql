package config

import (
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	AppHost     string
	Port        string
	DatabaseUrl string
	Tracing     bool
}

func New() Config {
	tracing, _ := strconv.ParseBool(os.Getenv("TRACING"))

	return Config{
		AppHost:     os.Getenv("APP_HOST"),
		Port:        os.Getenv("PORT"),
		DatabaseUrl: os.Getenv("DATABASE_URL"),
		Tracing:     tracing,
	}
}

func (c Config) DatabaseName() string {
	u, err := url.Parse(c.DatabaseUrl)

	if err != nil {
		panic(err)
	}

	return strings.Replace(u.Path, "/", "", 1)
}
