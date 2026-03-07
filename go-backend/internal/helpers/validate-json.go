package helpers

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)


func ValidateJSON[T any](body io.ReadCloser, validation *validator.Validate) (T, error) {
	var v T

	if err := json.NewDecoder(body).Decode(&v); err != nil {
		return  v, err
	}

	if err := validation.Struct(v); err != nil {
		return v, err
	}

	return v, nil
}