package controller

import (
	"eva/internal/controller/user"
	"eva/internal/repository"
)

type Controllers struct {
	userController *user.UserController
}

func NewControllers(repository *repository.Repository) *Controllers {
	return &Controllers{
		userController: ,
	}
}
