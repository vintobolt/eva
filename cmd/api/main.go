package main

import (
	"eva/internal/app"
	"eva/internal/config"
	"eva/internal/controllers"
	"eva/internal/repository"
	"eva/pkg/logging"
	"eva/pkg/postgresql"
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

	// Init postgresql connection pooler
	pgpool := postgresql.NewPgxPool(cfg.GetPostgresConnectionString(), &logger)

	// Init repo
	repo := repository.NewRepositories(pgpool, &logger)
	controllers := controllers.NewControllers(repo)
	app, err := app.NewApp(cfg, &logger, controllers)
	if err != nil {
		logger.Fatal(err)
	}
	app.Run()
}
