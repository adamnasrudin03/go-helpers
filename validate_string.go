package help

import (
	"net/mail"
	"regexp"
	"strconv"
	"strings"

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

// IsPhoneNumberId checks if the given string is a valid Indonesian phone number.
//
// Parameters:
// - input: The string to be validated.
//
// Returns:
// - A boolean indicating whether the string is a valid phone number or not.
//
// The function first removes any whitespace or dashes from the input string.
// Then it trims any non-digit characters from the input string, keeping only the
// digit characters and the "+" character, which is used to indicate the country
// code.
// The trimmed string is then matched against a regular expression pattern.
// If the trimmed string matches the pattern, the function returns true, otherwise
// it returns false.
func IsPhoneNumberId(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, "-", "")
	// The pattern to match the phone number. It should start with "+" or "62" or "0", followed by "8" and then 9 digits.
	pattern := `^(\+62|62|0)(8)\d{9}$`

	// Trim the input string to remove any non-digit characters. This is done to handle cases where the input string contains
	// white spaces or any other non-digit characters.
	phoneNumber := strings.TrimFunc(input, func(r rune) bool {
		// Only keep digit characters and the "+" character, which is used to indicate the country code.
		return !('0' <= r && r <= '9' || r == '+')
	})

	// Check if the phone number matches the pattern using the regexp.MatchString function.
	// If there is an error during the matching process, return false.
	match, err := regexp.MatchString(pattern, phoneNumber)
	if err != nil {
		return false
	}

	// Return the match result, indicating whether the phone number is valid or not.
	return match
}
