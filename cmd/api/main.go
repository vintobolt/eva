package main

import (
	"eva/internal/app"
	"eva/internal/config"
	"eva/pkg/logging"
	"log"
)

func main() {
	log.Println("EVA server starting")

	log.Println("Configuring EVA server...")
	cfg := config.GetConfig()
	log.Println("EVA server configured.")

	log.Println("Setting up logger...")
	logger := logging.GetLogger()
	log.Println("Logger set up.")

	app, err := app.NewApp(cfg, &logger)
	if err != nil {
		logger.Fatal(err)
	}

	app.Run()
}
