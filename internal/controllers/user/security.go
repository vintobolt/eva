package user

import (
	"eva/internal/config"
	"eva/internal/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var whiteListPaths = []string{
	"/api",
	"/api/*",
	"/api/v1/login",
	"/api/v1/signup",
}

// change default error message
func init() {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "Unauthorized"
}

func WebSecurityConfig(e *echo.Echo) {
	gc := config.GetConfig()
	config := middleware.JWTConfig{
		Claims:     &models.JwtCustomClaims{},
		SigningKey: []byte(gc.Server.JWTSecret),
		Skipper:    skipAuth,
	}
	e.Use(middleware.JWTWithConfig(config))
}

func skipAuth(e echo.Context) bool {
	// Skip authentication for and signup login requests
	for _, path := range whiteListPaths {
		if path == e.Path() {
			return true
		}
	}
	return false
}
