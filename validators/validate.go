package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func New() {
	Validate = validator.New()
}

func structValidator(err error) (bool, validator.ValidationErrors) {
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return true, nil
		}

		return true, err.(validator.ValidationErrors)
	}

	return false, nil
}
