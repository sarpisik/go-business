package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/sarpisik/go-business/models"
)

func LoginValidator(u *models.User) (bool, validator.ValidationErrors) {
	return structValidator(Validate.Struct(u))
}
