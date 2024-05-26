//go:build !test

package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

// Config Common config struct
type Config struct {
	Service Service
	Auth    Auth
}

// Service struct for storage this server config variables
type Service struct {
	Port string `env:"GATEWAY_SERVICE_PORT"`
}

// Auth struct for storage auth-service config variables
type Auth struct {
	Host string `env:"AUTH_SERVICE_HOST"`
	Port string `env:"AUTH_SERVICE_PORT"`
}

func MustLoad() *Config {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatalf("Can not read env variables: %s", err)
	}
	return cfg
}
