package v1

import (
	"log"
	"net/http"
	"strings"

	help "github.com/adamnasrudin03/go-helpers"
	"github.com/adamnasrudin03/go-helpers/validators"
	"github.com/go-playground/validator/v10"
)

// StatusMapping maps HTTP status code to a descriptive string.
// It returns the descriptive string which can be used as the 'status' field in ResponseDefault.
//
// Parameters:
// - statusCode: The HTTP status code to be mapped.
//
// Returns:
// - string: The descriptive string for the given HTTP status code.
//
// The function first defines a map to store custom status code mappings.
// It then checks if the given status code has a custom mapping.
// If it does, it returns the corresponding descriptive string.
// If not, it uses the default status text for the given status code.
// It then trims any leading or trailing white spaces from the status string.
// If the status string is still empty, it maps the status code to a default status code
// and returns the corresponding descriptive string.
func StatusMapping(statusCode int) string {
	// Define a map to store custom status code mappings.
	mappingsCustomStatus := map[int]string{
		http.StatusOK: "Success", // Map status code 200 to "Success".
	}

	// Get the status string from the custom mappings.
	status := mappingsCustomStatus[statusCode]

	// If the status string is still empty, use the default status text for the given status code.
	if status == "" {
		status = http.StatusText(statusCode)
	}

	// Trim any leading or trailing white spaces from the status string.
	status = strings.TrimSpace(status)

	// If the status string is empty, map the status code to a default status code
	// and return the corresponding descriptive string.
	if status == "" {
		statusCode = StatusErrorMapping(statusCode)
		status = mappingsCustomStatus[statusCode]
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

	// Translate the error messages from English to Indonesian and store the translated messages in msgIdn.
	msgIdn, errTranslate := help.Translate(help.Auto, help.LangID, msgEnUs)
	if errTranslate != nil {
		// If there is an error during translation, log the error and set msgIdn equal to msgEnUs.
		msgIdn = msgEnUs
		log.Printf("Translate Text %v to %v error: %v \n", help.Auto, help.LangID, errTranslate)
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
		msgEnUs.WriteString(validators.FormatErrorValidatorSingle(e))
	}
	return msgEnUs.String()
}
