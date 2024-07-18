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
// It takes a validator.FieldError object as input and returns a string representing the formatted error message.
// The function extracts relevant information from the error object, including the validation tag, field name, validation parameter, and field type.
// It then determines the appropriate error message based on the validation tag and constructs the message accordingly.
// Additionally, if the error is related to a string field and a parameter is provided, the function appends "character" to the error message.
//
// Parameters:
// - e: A validator.FieldError object containing information about the specific validation error.
//
// Returns:
// - A string containing the formatted error message based on the error details.
func FormatErrorValidatorSingle(e validator.FieldError) string {
	// Extract relevant information from the error object.
	tag := e.Tag()               // Obtain the validation tag associated with the error.
	field := e.Field()           // Retrieve the name of the field where the validation error occurred.
	param := e.Param()           // Get the validation parameter provided in the error.
	fieldType := e.Type().Name() // Determine the type of the field where the error occurred.

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
		msg = fmt.Sprintf("%v must be greater than %v", field, param)
	case "lt":
		msg = fmt.Sprintf("%v must be less than %v", field, param)
	case "eq":
		msg = fmt.Sprintf("%v must be equal to %v", field, param)
	case "eq_ignore_case":
		msg = fmt.Sprintf("%v must be equal to %v (ignoring case)", field, param)
	default:
		msg = fmt.Sprintf("%v is %v %v", field, tag, param)
	}

	// Append "character" if the error is for a string field and a parameter is provided.
	if param != "" && fieldType == "string" {
		msg += " character"
	}

	return msg
}
