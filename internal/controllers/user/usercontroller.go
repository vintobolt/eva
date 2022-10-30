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

// @Summary getting token
// @Description
// @Tags users
// @Accept json
// @Router /users/signup [post]
func (c *UserController) SignIn(ec echo.Context) error {
	signInData := models.SignIn{}
	if err := utils.BindAndValidate(ec, &signInData); err != nil {
		return err
	}
	user, valid := c.ValidateCredentials(signInData.Username, signInData.Password)
	if !valid {
		return exception.UnauthorizedException()
	}
	fmt.Printf("%+v\n", user)
	return nil
}

// @Summary Create a user
// @Description Take json and create an inactive user
// @Tags users
// @Accept  json
// @Success 200 {object} string "ok"
// @Router /users/signup [post]
func (c *UserController) SignUp(ec echo.Context) error {
	signUpData := models.SignUp{}
	if err := ec.Bind(signUpData); err != nil {
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
	fmt.Printf("%+v\n", signUpData)
	return ec.JSON(200, "ok")
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
