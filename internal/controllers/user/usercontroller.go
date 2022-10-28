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

// swagger:route POST /v1/users:login users login
func (c *UserController) GetExistUser(ec echo.Context) error {
	login := ec.QueryParam("login")
	user, err := c.userRepo.GetExistUser(login)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("user: %v", user)
	return ec.JSON(200, user)
}
