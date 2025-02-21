package main

import (
	"backend/internal/app"
)

func main() {
	application := app.New()

	if err := application.Run(); err != nil {
		panic(err)
	}
}
