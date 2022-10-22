package app

import (
	"context"
	"eva/internal/config"
	"eva/pkg/logging"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/swaggo/echo-swagger/example/docs"
)

type Appv interface {
	Release()
}

type App struct {
	cfg    *config.Config
	logger *logging.Logger
	e      *echo.Echo
}

func NewApp(config *config.Config, logger *logging.Logger) (App, error) {
	e := echo.New()
	logger.Info("swagger doc initializing")
	configureSwagger(e)
	configureMiddlewares(e, logger)
	configureCORS(e, logger)
	configureTimeouts(config, e)
	return App{cfg: config, logger: logger, e: e}, nil
}

func (a *App) Run() {
	//a.logger.Fatal(a.e.Start(a.cfg.GetServeString()))
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

func configureTimeouts(cfg *config.Config, e *echo.Echo) {
	e.Server.ReadTimeout = time.Duration(cfg.Server.ReadTimeout)
	e.Server.WriteTimeout = time.Duration(cfg.Server.WriteTimeout)
}

func configureSwagger(e *echo.Echo) {
	e.GET("/swagger", func(c echo.Context) error {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
		return nil
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

func configureMiddlewares(e *echo.Echo, logger *logging.Logger) {
	e.Use(middleware.Recover())
	logger.Info("Recover middleware used.")    // middleware for wrapping panics in chain
	e.Use(echozap.ZapLogger(logger.Desugar())) // using echozap instead default logger
	logger.Info("Used zap logger instead default.")
}

func configureCORS(e *echo.Echo, logger *logging.Logger) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))
	logger.Info("CORS configured.")
}
