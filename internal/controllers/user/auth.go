package user

import (
	"eva/internal/models"
	"eva/pkg/utils"
)

func (c *UserController) ValidateCredentials(username, password string) (*models.User, bool) {
	//user, err := c.userRepo.FindByEmail(username)
	user, err := c.userRepo.GetExistUser(username)
	if err != nil || utils.VerifyPassword(user.Passwd, password) != nil {
		return nil, false
	}
	return &user, true
}
