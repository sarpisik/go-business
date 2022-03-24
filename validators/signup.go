package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/sarpisik/go-business/models"
)

type SignUpUser struct {
	models.User
	ConfirmPassword string `validate:"required,eqfield=Password"`
}

func UserValidator(u *SignUpUser) (bool, validator.ValidationErrors) {
	return structValidator(Validate.Struct(u))
}
