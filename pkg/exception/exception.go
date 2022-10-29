package exception

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ResourceNotFoundException(resourceName string, fieldName string, fieldValue string) error {
	msg := fmt.Sprintf("Resource %s not found for field %s with value %s", resourceName, fieldName, fieldValue)
	return echo.NewHTTPError(http.StatusNotFound, msg)
}

func BadRequestException(msg string) error {
	return echo.NewHTTPError(http.StatusBadRequest, msg)
}

func CinflictException(resourceName, fieldName, fieldValue string) error {
	msg := fmt.Sprintf("%s with %s : %s already exist", resourceName, fieldName, fieldValue)
	return echo.NewHTTPError(http.StatusConflict, msg)
}

func UnauthorizedException() error {
	return echo.ErrUnauthorized
}
