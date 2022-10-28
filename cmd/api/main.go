package main

import (
	"eva/internal/app"
	"eva/internal/config"
	"eva/internal/controllers"
	"eva/internal/repository"
	"eva/pkg/logging"
	"eva/pkg/postgresql"
	"log"
	//_ "eva/docs"
)

// @title Swagger EVA API
// @version 0.1
// @description EVA API documentation
// @termsOfService http://swagger.io/terms/

// @contact.name vintobolt
// @contact.url vintobolt
// @contact.email vintobolt@protonmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8089
// @BasePath /api/v1/

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
