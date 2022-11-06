package app

import (
	"context"
	"eva/internal/config"
	"eva/internal/controllers"
	"eva/internal/routes"
	"eva/pkg/logging"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

/*
type Appv interface {
	Release()
}
*/

type App struct {
	cfg         *config.Config
	logger      *logging.Logger
	e           *echo.Echo
	controllers *controllers.Controllers
}

func NewApp(config *config.Config, logger *logging.Logger, controllers *controllers.Controllers) (App, error) {
	e := echo.New()
	logger.Info("swagger doc initializing")
	configureValidator(e)
	configureSwagger(e)
	configureMiddlewares(e, logger)
	configureCORS(e, logger)
	configureTimeouts(config, e)
	configureHealthCheck(e, logger)
	configureSecurity(e)
	routes.GetUserApiRoutes(e, *controllers)
	return App{cfg: config, logger: logger, e: e, controllers: controllers}, nil
}

func (a *App) Run() {
	go func() {
		if err := a.e.Start(a.cfg.GetServeString()); err != nil && err != http.ErrServerClosed {
			a.e.Logger.Fatal("Shutting down the server.")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(a.cfg.Server.ShutdownTimeout))
	defer cancel()
	if err := a.e.Shutdown(ctx); err != nil {
		a.e.Logger.Fatal(err)
	}
}
