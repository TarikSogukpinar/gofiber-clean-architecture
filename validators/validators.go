package validators

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(user interface{}) error {
	return validate.Struct(user)
}
