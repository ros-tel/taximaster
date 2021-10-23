package validator

import v "github.com/go-playground/validator/v10"

func Validate(s interface{}) error {
	validate := v.New()
	return validate.Struct(s)
}
