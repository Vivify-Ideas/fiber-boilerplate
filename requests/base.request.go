package requests

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse map[string]string

func FormatValidationResponse(err error) ErrorResponse {
	errors := make(map[string]string)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			rule := "[" + err.Tag()
			if err.Param() != "" {
				rule = rule + ": " + err.Param()
			}
			rule = rule + "]"

			errors[field] = "Validation failed. Not fulfilled rule " + rule
		}
		return errors
	}
	return nil
}

func Validate(data interface{}) ErrorResponse {
	validate := validator.New()
	err := validate.Struct(data)
	return FormatValidationResponse(err)
}
