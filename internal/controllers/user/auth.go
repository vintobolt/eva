package user

import (
	"eva/internal/models"
	"eva/pkg/utils"
)

func (c *UserController) ValidateCredentials(username, password string) (*models.SignIn, bool) {
	//user, err := c.userRepo.FindByEmail(username)
	user, err := c.userRepo.FindCredsWithUsername(username)
	if err != nil || utils.VerifyPassword(user.Password, password) != nil {
		return nil, false
	}
	return user, true
}
