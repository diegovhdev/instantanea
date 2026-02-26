package helpers

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(v *validator.Validate, s any) map[string]string {
	err := v.Struct(s)
	if err == nil {
		return nil
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return map[string]string{
			"error": "invalid request",
		}
	}

	errors := make(map[string]string)

	for _, fieldErr := range validationErrors {
		field := strings.ToLower(fieldErr.Field())

		errors[field] = buildErrorMessage(fieldErr)
	}

	return errors
}

func buildErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {

	case "required":
		return "is required"

	case "email":
		return "must be a valid email"

	case "min":
		return "must be at least " + fe.Param() + " characters"

	case "max":
		return "must be at most " + fe.Param() + " characters"

	case "gte":
		return "must be greater than or equal to " + fe.Param()

	case "lte":
		return "must be less than or equal to " + fe.Param()

	case "eqfield":
		return "must match " + fe.Param()

	default:
		return "is invalid"
	}
}