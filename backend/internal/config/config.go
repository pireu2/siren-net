package config

import (
	"os"
	"time"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
	TokenExpiry time.Duration
	SDUrl       string
}

func Load() Config {
	return Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		TokenExpiry: time.Second * 10,
		SDUrl:       "https://dd2e43242112719bfa.gradio.live",
	}
}