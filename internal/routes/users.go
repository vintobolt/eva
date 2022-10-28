package routes

import (
	"eva/internal/controllers"

	"github.com/labstack/echo/v4"
)

func GetUserApiRoutes(e *echo.Echo, c controllers.Controllers) {
	v1 := e.Group("/api/v1")
	{
		v1.GET("/users/:login", c.UserController.GetExistUser)
		v1.POST("/users/signin", c.UserController.SignIn)
		v1.POST("/users/signup", c.UserController.SignUp)
	}
}
