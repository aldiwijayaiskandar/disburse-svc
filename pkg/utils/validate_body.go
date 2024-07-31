package utils

import (
	"github.com/gookit/validate"
)

func ValidateBody(body interface{}) error {
	// validate dto
	valid := validate.Struct(body)

	if !valid.Validate() {
		return valid.Errors.OneError()
	}

	return nil
}
