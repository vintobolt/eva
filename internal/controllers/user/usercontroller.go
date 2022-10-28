package user

import (
	"eva/internal/repository/users"
	"fmt"

	"github.com/labstack/echo/v4"
)

// declaring the repository interface in the controller package
// allows us to easily swap out the actual implementation, enforcing loose coupling
/*
type repository interface {
	GetExistUser(login string) (models.User, error)
	CreateUser()
	DeleteUser()
}
*/
// Controller contains the service, which contains database-related logic
// as an injectable dependency, allowing us to decouple business logic from db package
type UserController struct {
	userRepo users.UserRepository
}

// Initialize the user controller.
func NewController(userRepository users.UserRepository) *UserController {
	return &UserController{
		userRepo: userRepository,
	}
}

// GetUser godoc
// @Summary Get a user
// @Description Get a user item
// @Tags users
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json, xml)
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func (c *UserController) GetExistUser(ec echo.Context) error {
	login := ec.Param("login")
	user, err := c.userRepo.GetExistUser(login)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("user: %v", user)
	return ec.JSON(200, user)
}

func (c *UserController) SignIn(ec echo.Context) error {
	fmt.Println(ec.Request().Header.Get("Authorization"))
	login := ec.FormValue("login")
	password := ec.FormValue("password")
	fmt.Println("heheh", login, password)
	return ec.JSON(200, "ok")
	//return nil
}

func (c *UserController) SignUp(ec echo.Context) error {
	return nil
}

func (c *UserController) RefreshToken(ec echo.Context) error {
	return nil
}
