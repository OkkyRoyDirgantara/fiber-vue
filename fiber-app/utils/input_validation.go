package utils

import "github.com/go-playground/validator/v10"

func ValidationInput(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}
