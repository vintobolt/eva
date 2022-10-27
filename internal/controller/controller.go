package controller

import (
	"eva/internal/controller/user"
	"eva/internal/repository"
)

// Controllers contains all the controllers
type Controllers struct {
	userController *user.UserController
}

// Constructor return a new Controllers
func NewControllers(repositories *repository.Repositories) *Controllers {
	return &Controllers{
		userController: user.NewController(repositories.UserRepo),
	}
}
