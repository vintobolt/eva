package app

import (
	"eva/internal/config"
	"eva/pkg/logging"
	"eva/pkg/utils"
	"net/http"
	"time"

	_ "eva/docs"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// set http timeouts.
func configureTimeouts(cfg *config.Config, e *echo.Echo) {
	e.Server.ReadTimeout = time.Duration(cfg.Server.ReadTimeout) * time.Second
	e.Server.WriteTimeout = time.Duration(cfg.Server.WriteTimeout) * time.Second
}

// enable validation
func configureValidator(e *echo.Echo) {
	e.Validator = utils.NewValidatorUtil()
}

// serve swagger files
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

// healthcheck
func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func configureHealthCheck(e *echo.Echo, logger *logging.Logger) {
	e.GET("/health", healthCheck)
	logger.Info("Health check configured.")
}
