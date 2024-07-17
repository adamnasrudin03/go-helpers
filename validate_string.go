package help

import (
	"net/mail"
	"strconv"

	"github.com/google/uuid"
)

// IsEmail checks if the given email is in a valid format.
//
// Parameters:
// - email: The email address to be validated.
//
// Returns:
// - A boolean indicating whether the email is valid or not.
func IsEmail(email string) bool {
	// Parse the email address using the mail.ParseAddress function.
	// This function returns the parsed email address and an error.
	_, err := mail.ParseAddress(email)

	// Check if the error is nil, indicating that the email is valid.
	return err == nil
}

// IsUUID checks if the given string is a valid UUID.
//
// Parameters:
// - input: The string to be validated.
//
// Returns:
// - A boolean indicating whether the string is a valid UUID or not.
func IsUUID(input string) bool {
	// Parse the given string using the uuid.Parse function.
	// This function returns a UUID and an error.
	_, err := uuid.Parse(input)

	// Check if the error is nil, indicating that the string is a valid UUID.
	return err == nil
}

// IsNumber checks if the given string is a valid number.
//
// Parameters:
// - input: The string to be validated.
//
// Returns:
// - A boolean indicating whether the string is a valid number or not.
func IsNumber(input string) bool {
	// Convert the string to an integer using strconv.Atoi.
	// This function returns the integer value and an error.
	_, err := strconv.Atoi(input)

	// Check if the error is nil, indicating that the string is a valid number.
	return err == nil
}
