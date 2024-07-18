package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// FormatErrorValidator formats multiple validation error messages.
//
// It takes a slice of validator.ValidationErrors and returns a slice of strings,
// where each string is a formatted error message.
func FormatErrorValidator(errs validator.ValidationErrors) []string {
	// Create a slice to hold the formatted error messages.
	var msgEnUs []string

	// Loop through each error and append the formatted error message to the slice.
	for _, e := range errs {
		msgEnUs = append(msgEnUs, FormatErrorValidatorSingle(e))
	}

	// Return the slice of formatted error messages.
	return msgEnUs
}

// FormatErrorValidatorSingle formats a single validation error message.
//
// It takes a validator.FieldError and returns a string.
// The function extracts relevant information from the error,
// determines the error message based on the validation tag,
// and appends "character" if the error is for a string field and a parameter is provided.
// The function returns the formatted error message.
func FormatErrorValidatorSingle(e validator.FieldError) string {
	// Extract relevant information from the error.
	tag := e.Tag()     // Validation tag.
	field := e.Field() // Field name.
	param := e.Param() // Validation parameter.

	// Determine the error message based on the validation tag.
	var msg string
	switch tag {
	case "required":
		msg = fmt.Sprintf("%s is a required field", field)
	case "email":
		msg = fmt.Sprintf("%v must be an email address", field)
	case "e164":
		msg = fmt.Sprintf("%v must be a valid phone number in E.164 format", field)
	case "gte":
		msg = fmt.Sprintf("%v must be greater than or equal to %v", field, param)
	case "lte":
		msg = fmt.Sprintf("%v must be less than or equal to %v", field, param)
	case "gt":
		msg = fmt.Sprintf("%v must be greater than to %v", field, param)
	case "lt":
		msg = fmt.Sprintf("%v must be less than to %v", field, param)
	case "eq":
		msg = fmt.Sprintf("%v must be equals to %v", field, param)
	case "eq_ignore_case":
		msg = fmt.Sprintf("%v must be equals ignoring case to %v", field, param)
	default:
		msg = fmt.Sprintf("%v is %v %v", field, tag, param)
	}

	// Append "character" if the error is for a string field and a parameter is provided.
	if param != "" && e.Type().Name() == "string" {
		msg += " character"
	}

	return msg
}
