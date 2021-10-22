package common_api

import "github.com/go-playground/validator/v10"

func validate(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}
