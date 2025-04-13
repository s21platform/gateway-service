//go:build !test

package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type key string

const (
	KeyMetrics  = key("metrics")
	KeyUUID     = key("uuid")
	KeyUsername = key("username")
	KeyLogger   = key("logger")
)

// Config Common config struct
type Config struct {
	Service      Service
	Auth         Auth
	User         User
	Avatar       Avatar
	Friends      Friends
	Option       Option
	Metrics      Metrics
	Platform     Platform
	Notification Notification
	Society      Society
	Logger       Logger
	Search       Search
	Chat         Chat
	Advert       Advert
	Feed         Feed
}

// Service struct for storage this server config variables
type Service struct {
	Port string `env:"GATEWAY_SERVICE_PORT"`
	Name string `env:"GATEWAY_SERVICE_NAME"`
}

// Auth struct for storage auth-service config variables
type Auth struct {
	Host string `env:"AUTH_SERVICE_HOST"`
	Port string `env:"AUTH_SERVICE_PORT"`
}

type User struct {
	Host string `env:"USER_SERVICE_HOST"`
	Port string `env:"USER_SERVICE_PORT"`
}

type Avatar struct {
	Host string `env:"AVATAR_SERVICE_HOST"`
	Port string `env:"AVATAR_SERVICE_PORT"`
}

type Option struct {
	Host string `env:"OPTIONHUB_SERVICE_HOST"`
	Port string `env:"OPTIONHUB_SERVICE_PORT"`
}

type Friends struct {
	Host string `env:"FRIENDS_SERVICE_HOST"`
	Port string `env:"FRIENDS_SERVICE_PORT"`
}

type Notification struct {
	Host string `env:"NOTIFICATION_SERVICE_HOST"`
	Port string `env:"NOTIFICATION_SERVICE_PORT"`
}

type Metrics struct {
	Host string `env:"GRAFANA_HOST"`
	Port int    `env:"GRAFANA_PORT"`
}

type Platform struct {
	Secret string `env:"SECRET_KEY"`
	Env    string `env:"ENV"`
}

type Society struct {
	Host string `env:"SOCIETY_SERVICE_HOST"`
	Port string `env:"SOCIETY_SERVICE_PORT"`
}

type Logger struct {
	Host string `env:"LOGGER_SERVICE_HOST"`
	Port string `env:"LOGGER_SERVICE_PORT"`
}

type Search struct {
	Host string `env:"SEARCH_SERVICE_HOST"`
	Port string `env:"SEARCH_SERVICE_PORT"`
}

type Chat struct {
	Host string `env:"CHAT_SERVICE_HOST"`
	Port string `env:"CHAT_SERVICE_PORT"`
}

type Advert struct {
	Host string `env:"ADVERT_SERVICE_HOST"`
	Port string `env:"ADVERT_SERVICE_PORT"`
}

type Feed struct {
	Host string `env:"FEED_SERVICE_HOST"`
	Port string `env:"FEED_SERVICE_PORT"`
}

func MustLoad() *Config {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatalf("Can not read env variables: %s", err)
	}
	return cfg
}
