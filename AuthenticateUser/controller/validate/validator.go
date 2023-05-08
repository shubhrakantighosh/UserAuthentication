package validate

import (
	"AuthenticateUser/errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"time"
	"unicode"
)

func containsUppercase(s string) bool {
	for _, value := range s {
		if unicode.IsUpper(value) {
			return true
		}
	}
	return false
}

func containsDigit(s string) bool {
	for _, value := range s {
		if unicode.IsDigit(value) {
			return true
		}
	}
	return false
}

func containsSpace(s string) bool {
	for _, value := range s {
		if unicode.IsSpace(value) {
			return false
		}
	}
	return true
}

var validate = validator.New()

func Get(request interface{}) errors.CustomError {
	var cusErr errors.CustomError

	validate.RegisterValidation("uppercase", func(fl validator.FieldLevel) bool {
		return containsUppercase(fl.Field().String())
	})

	validate.RegisterValidation("number", func(fl validator.FieldLevel) bool {
		return containsDigit(fl.Field().String())
	})

	validate.RegisterValidation("space", func(fl validator.FieldLevel) bool {
		return containsSpace(fl.Field().String())
	})

	if err := validate.Struct(request); err != nil {
		errors := err.(validator.ValidationErrors)
		messages := make(map[string]string, 0)
		for _, e := range errors {
			switch e.Tag() {
			case "required":
				messages[e.Field()] = " is required"
			case "min":
				messages[e.Field()] = fmt.Sprintf(" should have at least %s characters", e.Param())
			case "max":
				messages[e.Field()] = fmt.Sprintf(" should have at most %s characters", e.Param())
			case "len":
				messages[e.Field()] = fmt.Sprintf(" must be exactly %s ", e.Param())
			case "email":
				messages[e.Field()] = " is not a valid email"
			case "url":
				messages[e.Field()] = " is url is not valid"
			case "oneof":
				messages[e.Field()] = fmt.Sprintf(" must be one of %s", e.Param())
			case "uuid":
				messages[e.Field()] = " is not a valid uuid"
			case "latitude":
				messages[e.Field()] = " is not valid latitude"
			case "longitude":
				messages[e.Field()] = " is not valid longitude"
			case "numeric":
				messages[e.Field()] = " only numeric values are allowed"
			case "alpha":
				messages[e.Field()] = " only alphabets are allowed"
			case "alphanum":
				messages[e.Field()] = " only numeric and alphabets values are allowed"
			case "uppercase":
				messages[e.Field()] = " one minimum uppercase"
			case "number":
				messages[e.Field()] = " one minimum number"
			case "containsany":
				messages[e.Field()] = fmt.Sprintf(" a minimum of 1 special character  %s", e.Param())
			case "space":
				messages[e.Field()] = " spaces are not allowed"
			default:
				messages[e.Field()] = "is required"
			}
		}

		cusErr.Time = time.Now()
		cusErr.ErrorExist = true
		cusErr.UserMessage = "Haanji errors toh hogi, tumne data i dala hain galat"
		cusErr.StructErrors = messages
	}
	return cusErr
}
