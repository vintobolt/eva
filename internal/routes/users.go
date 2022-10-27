package routes

import (
	"eva/internal/controllers"

	"github.com/labstack/echo/v4"
)

func GetUserApiRoutes(e *echo.Echo, c controllers.Controllers) {
	e.GET("/users", c.UserController.GetExistUser)
}
