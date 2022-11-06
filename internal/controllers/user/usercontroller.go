package user

import (
	"eva/internal/models"
	"eva/internal/repository/users"
	"eva/pkg/exception"
	"eva/pkg/utils"
	"fmt"

	"github.com/labstack/echo/v4"
)

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
// @Accept json
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json)
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /users/{login} [get]
// @Security ApiKeyAuth
func (c *UserController) GetExistUser(ec echo.Context) error {
	login := ec.Param("login")
	user, err := c.userRepo.GetExistUser(login)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("user: %v", user)
	return ec.JSON(200, user)
}

// TODO: add logger
func (c *UserController) SignIn(ec echo.Context) error {
	payload := models.SignIn{}
	if err := utils.BindAndValidate(ec, &payload); err != nil {
		return err
	}
	user, valid := c.ValidateCredentials(payload.Username, payload.Password)
	if !valid {
		return exception.UnauthorizedException()
	}

	jwt, err := utils.GenerateJwtToken(user)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("%+v\n", user)
	return ec.JSON(200, models.Token{Token: jwt})
}

// SignUp godoc
// @Summary Create an inactive user
// @Description Create a new user item
// @Tags users
// @Accept json
// @Produce json
// @Param mediaType query string false "mediaType" Enums(json)
// @Param user body models.SignUp true "New User"
// @Success 200 {object} models.SignUp
// @Failure 400 {object} handler.APIError
// @Failure 409 {object} handler.APIError
// @Failure 500 {object} handler.APIError
// @Router /signup [post]
func (c *UserController) SignUp(ec echo.Context) error {
	signUpData := models.SignUp{}
	if err := utils.BindAndValidate(ec, &signUpData); err != nil {
		fmt.Println("HERE::\t", err)
		return err
	}
	_, err := c.userRepo.FindCredsWithUsername(signUpData.Username)
	if err == nil {
		return exception.CinflictException("User", "User", signUpData.Username)
	}
	err = beforeSave(&signUpData)
	if err != nil {
		return err
	}
	err = c.userRepo.CreateUser(signUpData)
	if err != nil {
		return err
	}

	return ec.JSON(200, signUpData)
}

func (c *UserController) RefreshToken(ec echo.Context) error {
	return nil
}

func beforeSave(user *models.SignUp) (err error) {
	hashedPassword, err := utils.EncryptPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}
