package utils

import (
	"eva/pkg/exception"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ValidatorUtil struct {
	validator *validator.Validate
}

func NewValidatorUtil() *ValidatorUtil {
	return &ValidatorUtil{validator: validator.New()}
}

func (v *ValidatorUtil) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func BindAndValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return exception.BadRequestException(err.Error())
	}
	if err := c.Validate(i); err != nil {
		return exception.BadRequestException(err.Error())
	}
	return nil
}
