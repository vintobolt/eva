package user

import (
	"eva/internal/models"
	"eva/internal/repository/users"
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

func (c *UserController) GetExistUser(login string) (models.User, error) {
	return c.userRepo.GetExistUser(login)
}
