package controllers

import (
	"eva/internal/controllers/user"
	"eva/internal/repository"
)

// Controllers contains all the controllers
type Controllers struct {
	UserController *user.UserController
}

// Constructor return a new Controllers
func NewControllers(repositories *repository.Repositories) *Controllers {
	return &Controllers{
		UserController: user.NewController(repositories.UserRepo),
	}
}
