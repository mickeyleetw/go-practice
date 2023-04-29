package core

import (
	"github.com/dlclark/regexp2"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

func ValidateAccountPassword(fl validator.FieldLevel) bool {
	name, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	pattern := `^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])[a-zA-z0-9]{8,32}$`
	regexp, _ := regexp2.Compile(pattern, 0)
	m, _ := regexp.FindStringMatch(name)
	return m != nil
}
