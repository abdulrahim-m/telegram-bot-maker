package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

var App Config

type Config struct {
	Port string `env:"PORT" envDefault:"3033"`
}

func Load() error {
	godotenv.Load()
	return env.Parse(&App)
}
