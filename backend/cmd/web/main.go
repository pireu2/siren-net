package main

import (
	"backend/internal/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load environment variables")
	}

	application := app.New()

	if err := application.Run(); err != nil {
		panic(err)
	}
}
