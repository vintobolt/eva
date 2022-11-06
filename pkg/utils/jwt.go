package utils

import (
	"eva/internal/config"
	"eva/internal/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateJwtToken(user *models.User) (string, error) {
	cfg := config.GetConfig()
	expTimeMs := cfg.Server.JWTExpirationMs
	exp := time.Now().Add(time.Millisecond * time.Duration(expTimeMs)).Unix()
	name := fmt.Sprintf("%s", user.Fullname)

	// Set custom claims
	claims := &models.JwtCustomClaims{
		user.ID,
		name,
		jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response
	jwt, err := token.SignedString([]byte(cfg.Server.JWTSecret))
	return jwt, err
}

func GetUserIdFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*models.JwtCustomClaims)
	return claims.ID
}
