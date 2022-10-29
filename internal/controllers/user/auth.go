package user

import (
	"eva/internal/models"
	"eva/pkg/utils"
)

func (c *UserController) ValidateCredentials(username, password string) (*models.User, bool) {
	user, err := c.userRepo.FindByEmail(username)
	if err != nil || utils.VerifyPassword(user.Password, password) != nil {
		return nil, false
	}
	return user, true
}

func (c *UserController) findByUsername() (models.SignIn, error) {

}
