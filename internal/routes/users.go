package routes

import (
	"eva/internal/controllers"

	"github.com/labstack/echo/v4"
)

func GetUserApiRoutes(e *echo.Echo, c controllers.Controllers) {
	v1 := e.Group("/v1")
	{
		v1.GET("/users/:login", c.UserController.GetExistUser)
	}
}
