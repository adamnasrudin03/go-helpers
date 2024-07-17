package help

import (
	"strings"
)

// CheckStringValue returns an empty string if the input pointer to a string is nil,
// otherwise it returns the value of the input pointer.
//
// Parameters:
// - input: A pointer to a string.
//
// Returns:
// - A string.
func CheckStringValue(input *string) string {
	// Check if the input pointer is nil.
	if input == nil {
		// If the input pointer is nil, return an empty string.
		return ""
	}
	// If the input pointer is not nil, return the value of the input pointer.
	return *input
}

// CheckStringValueToPointer returns a pointer to a string or nil if the input string is empty or only contains whitespace.
//
// Parameters:
// - input: A string.
//
// Returns:
// - A pointer to a string.
func CheckStringValueToPointer(input string) *string {
	// Trim the input string to remove any leading or trailing whitespace.
	input = strings.TrimSpace(input)

	// If the trimmed string is empty, return nil.
	if len(input) == 0 {
		return nil
	}

	// Otherwise, return a pointer to the trimmed string.
	return &input
}
