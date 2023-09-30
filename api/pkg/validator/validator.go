package validator

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(s interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validationError := validate.Struct(s)

	if validationError == nil {
		return nil
	}

	if _, ok := validationError.(*validator.InvalidValidationError); ok {
		return validationError
	}

	errors_map := make(map[string]string)
	for _, err := range validationError.(validator.ValidationErrors) {
		errors_map[err.Field()] = fmt.Sprintf("Failed on tag %s", err.Tag())
	}

	errorsString, err := json.Marshal(errors_map)
	if err != nil {
		return validationError
	}

	return errors.New(string(errorsString))
}
