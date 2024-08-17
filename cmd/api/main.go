package main

import (
	"log"

	"github.com/Uno-count/Task-Management/internal/config"
)

func main() {
	app := config.NewApp()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
