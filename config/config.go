package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	DatabaseUrl string `envconfig:"DATABASE_URL" split_words:"true" default:"postgres://localhost/dealgrok_dev?sslmode=disable"`
}

func Init() Config {
	var c Config
	err := envconfig.Process("dealgrok", &c)
	if err != nil {
		log.Fatalf("Unable to process environment: %s", err)
	}
	return c
}
