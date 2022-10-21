package main

import (
	"eva/internal/config"
	"eva/internal/repository"
	"eva/pkg/logging"
	"eva/pkg/postgresql"
	"fmt"
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

	pgpool := postgresql.NewPgxPool(cfg.GetPostgresConnectionString(), &logger)

	repo := repository.NewRepository(pgpool, &logger)
	userRole, err := repo.UserRepository.GetRoleByLogin("vintobot")
	if err != nil {
		logger.Error(err)
	}
	fmt.Printf("%+v", userRole)

	/*
		app, err := app.NewApp(cfg, &logger)
		if err != nil {
			logger.Fatal(err)
		}

		app.Run()*/
}
