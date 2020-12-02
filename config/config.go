package config

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

// Config is the main binary configuration.
type Config struct {
	ClientID     string
	ClientSecret string
	Email        string
	Password     string
	VIN          string
}

// FromEnv returns a Config from the environment variables.
func FromEnv() Config {
	var c Config
	err := envconfig.Process("gtg", &c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
