package v1

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	go_helpers_language "github.com/adamnasrudin03/go-helpers/language"
	"github.com/go-playground/validator/v10"
	"golang.org/x/text/language"
)

// StatusMapping maps HTTP status code to a descriptive string.
// It returns the descriptive string which can be used as the 'status' field in ResponseDefault.
//
// Parameters:
// - statusCode: The HTTP status code to be mapped.
//
// Returns:
// - string: The descriptive string for the given HTTP status code.
func StatusMapping(statusCode int) string {
	// Define a map to store custom status code mappings.
	mappingsCustomStatus := map[int]string{
		http.StatusOK: "Success", // Map status code 200 to "Success".
	}

	// Get the status string from the custom mappings.
	status := mappingsCustomStatus[statusCode]

	// If the status string is empty, map the status code to a default status code.
	if status == "" {
		statusCode = StatusErrorMapping(statusCode)
		status = mappingsCustomStatus[statusCode]
	}

	// Trim any leading or trailing white spaces from the status string.
	status = strings.TrimSpace(status)

	// If the status string is still empty, use the default status text for the given status code.
	if status == "" {
		status = http.StatusText(statusCode)
	}

	// Return the mapped status string.
	return status
}

// FormatValidationError is a function that formats validation errors during user input validation.
//
// It takes an error as input and returns an error.
// The function first checks if the error is of type validator.ValidationErrors.
// If it is, it formats the error messages and stores them in msgEnUs.
// The function then translates the error messages from English to Indonesian and stores the translated messages in msgIdn.
// If there is an error during translation, the function logs the error and sets msgIdn equal to msgEnUs.
// Finally, the function returns a new error of type *ResponseError with the translated error messages.
func FormatValidationError(err error) error {
	var (
		msgIdn  string // Holds the translated error messages
		msgEnUs string // Holds the English error messages
	)

	// Check if the error is of type validator.ValidationErrors.
	if errValidate := err.(validator.ValidationErrors); errValidate != nil {
		// Format the error messages and store them in msgEnUs.
		msgEnUs = formatMessageValidator(errValidate)
		msgEnUs = strings.TrimSpace(msgEnUs) + "."
	}

	// Define the language to translate to.
	langTo := language.Indonesian.String()
	// Define the language to translate from (auto detect).
	auto := go_helpers_language.Auto

	// Translate the error messages from English to Indonesian and store the translated messages in msgIdn.
	msgIdn, errTranslate := go_helpers_language.Translate(auto, langTo, msgEnUs)
	if errTranslate != nil {
		// If there is an error during translation, log the error and set msgIdn equal to msgEnUs.
		msgIdn = msgEnUs
		log.Printf("Translate Text %v to %v error: %v \n", auto, langTo, errTranslate)
	}

	// Return a new error of type *ResponseError with the translated error messages.
	return NewError(ErrValidation, NewResponseMultiLang(MultiLanguages{
		ID: msgIdn,
		EN: msgEnUs,
	}))
}

// formatMessageValidator formats multiple validation error messages.
//
// It takes a slice of validator.ValidationErrors and returns a single formatted string.
func formatMessageValidator(errs validator.ValidationErrors) string {
	var msgEnUs strings.Builder
	msgEnUs.Grow(len(errs) * 20) // rough estimate of the length of the joined string
	for i, e := range errs {
		if i > 0 {
			msgEnUs.WriteString(", ")
		}
		msgEnUs.WriteString(formatMessageValidatorSingle(e))
	}
	return msgEnUs.String()
}

// formatMessageValidatorSingle formats a single validation error message.
//
// It takes a validator.FieldError and returns a string.
func formatMessageValidatorSingle(e validator.FieldError) string {
	// Extract relevant information from the error.
	tag := e.Tag()
	field := e.Field()
	param := e.Param()

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
	default:
		msg = fmt.Sprintf("%v is %v %v", field, tag, param)
	}

	// Append "character" if the error is for a string field and a parameter is provided.
	if param != "" && e.Type().Name() == "string" {
		msg += " character"
	}

	return msg
}
