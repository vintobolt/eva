package app

import (
	"eva/internal/config"
	"eva/internal/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var whiteListPaths = []string{
	"/api",
	"/api/*",
	"/api/v1/users/signin",
	"/api/v1/users/signup",
	"/users/signin",
	"/users/signup",
	"/health",
}

func init() {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "Unauthorized"
}

func configureSecurity(e *echo.Echo) {
	config := middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte(config.GetConfig().Server.JWTSecret),
		Skipper:    skipAuth,
	}
	e.Use(middleware.JWTWithConfig(config))
}

func skipAuth(e echo.Context) bool {
	for _, path := range whiteListPaths {
		if path == e.Path() {
			return true
		}
	}
	return false
}
