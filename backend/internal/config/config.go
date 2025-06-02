package config

import (
	"os"
	"time"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
	TokenExpiry time.Duration
}

func Load() Config {
	return Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		TokenExpiry: time.Second * 3600,
	}
}
