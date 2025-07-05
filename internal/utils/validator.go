package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"unicode"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) Validate(data interface{}) map[string]string {
	errors := make(map[string]string)

	if errs := v.validator.Struct(data); errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			field := toSnakeCase(err.Field())
			errors[field] = getErrorMessage(err, field)
		}
	}

	return errors
}

func getErrorMessage(fe validator.FieldError, field string) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", field, fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters long", field, fe.Param())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters long", field, fe.Param())
	default:
		return fmt.Sprintf("%s is not valid", field)
	}
}

func toSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if unicode.IsUpper(r) {
			// Add underscore before uppercase letters (except the first character)
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func (v *Validator) IsValid(data interface{}) bool {
	return len(v.Validate(data)) == 0
}
