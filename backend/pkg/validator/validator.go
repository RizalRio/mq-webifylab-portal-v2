package validator

import (
	"fmt"
	"reflect"
	"strings"

	"webifylab-backend/pkg/response"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// Init menginisialisasi validator dengan custom rules
func Init() {
	validate = validator.New()

	// Custom naming: ubah field name ke snake_case untuk error message
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Custom validation rules bisa ditambahkan di sini
	// validate.RegisterValidation("custom_rule", customFunc)
}

// Validate memvalidasi struct dan mengembalikan error response
func Validate(s interface{}) []response.Error {
	if validate == nil {
		Init()
	}

	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	var errors []response.Error

	// Handle ValidationErrors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors = append(errors, response.Error{
				Field:   toSnakeCase(e.Field()),
				Message: formatErrorMessage(e),
			})
		}
	} else {
		// Handle invalid type errors
		errors = append(errors, response.Error{
			Message: err.Error(),
		})
	}

	return errors
}

// formatErrorMessage membuat pesan error yang user-friendly
func formatErrorMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email address", e.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", e.Field(), e.Param())
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", e.Field(), e.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of: %s", e.Field(), e.Param())
	case "url":
		return fmt.Sprintf("%s must be a valid URL", e.Field())
	case "numeric":
		return fmt.Sprintf("%s must be numeric", e.Field())
	case "alphanum":
		return fmt.Sprintf("%s must contain only letters and numbers", e.Field())
	case "uuid":
		return fmt.Sprintf("%s must be a valid UUID", e.Field())
	case "eqfield":
		return fmt.Sprintf("%s must match %s", e.Field(), e.Param())
	case "nefield":
		return fmt.Sprintf("%s must not be the same as %s", e.Field(), e.Param())
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", e.Field(), e.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", e.Field(), e.Param())
	case "lt":
		return fmt.Sprintf("%s must be less than %s", e.Field(), e.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", e.Field(), e.Param())
	case "file":
		return fmt.Sprintf("%s must be a valid file", e.Field())
	case "image":
		return fmt.Sprintf("%s must be a valid image", e.Field())
	default:
		return fmt.Sprintf("%s is invalid", e.Field())
	}
}

// toSnakeCase mengubah CamelCase ke snake_case
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteByte('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}