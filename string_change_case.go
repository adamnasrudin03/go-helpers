package help

import (
	"strings"
	"unicode"
)

// ToLower converts a given string to lower case.
//
// Parameters:
// - input: The string to be converted.
//
// Returns:
// - The lower case version of the input string..
//
// Example:
// - HELLO WORLD => hello world
//
// Check https://adamnasrudin.vercel.app/cheat-sheet/change-case-string-in-golang
func ToLower(input string) string {
	// Trim any leading or trailing whitespace from the input string.
	input = strings.TrimSpace(input)

	// Convert the input string to lower case using the ToLower function from the strings package.
	// The returned string is then trimmed to remove any leading or trailing whitespace.
	return strings.ToLower(input)
}

// ToUpper converts a given string to upper case.
// Parameters:
// - input: The string to be converted.
//
// Returns:
// - The upper case version of the input string.
//
// Example:
// - hello world => HELLO WORLD
//
// Check https://adamnasrudin.vercel.app/cheat-sheet/change-case-string-in-golang
func ToUpper(input string) string {
	// Trim any leading or trailing whitespace from the input string.
	input = strings.TrimSpace(input)

	// Convert the input string to upper case using the ToUpper function from the strings package.
	// The returned string is then trimmed to remove any leading or trailing whitespace.
	return strings.ToUpper(input)
}

// ToTitle converts a given string to title case.
//
// Parameters:
// - input: The string to be converted.
//
// Returns:
// - The title case version of the input string.
//
// Example:
// - hello world => Hello World
//
// Check https://adamnasrudin.vercel.app/cheat-sheet/change-case-string-in-golang
func ToTitle(input string) string {
	// Create an empty slice to store the output.
	var (
		output []rune
		isWord = true // Flag to indicate if we are currently in a word.
	)

	// Iterate over each character in the input string.
	for _, val := range input {
		// Check if we are currently in a word and the character is a letter.
		if isWord && unicode.IsLetter(val) {
			// If so, convert the character to upper case and append it to the output.
			output = append(output, unicode.ToUpper(val))

			// Set the flag to indicate that we are no longer in a word.
			isWord = false
		} else if !unicode.IsLetter(val) {
			// If the character is not a letter, set the flag to indicate that we are in a word.
			isWord = true

			// Append the character to the output.
			output = append(output, val)
		} else {
			// If we are not in a word and the character is a letter, simply append it to the output.
			output = append(output, val)
		}
	}

	// Trim any leading or trailing whitespace from the output and return it as a string.
	return strings.TrimSpace(string(output))
}

// ToSentenceCase converts a given string to sentence case.
//
// Parameters:
// - input: The string to be converted.
//
// Returns:
// - The sentence case version of the input string.
//
// Example:
// - hello world => Hello World
//
// Check https://adamnasrudin.vercel.app/cheat-sheet/change-case-string-in-golang
func ToSentenceCase(input string) string {
	// Convert the input string to lowercase.
	input = ToLower(input)

	// Check if the input string is empty.
	if len(input) <= 0 {
		return ""
	}

	// Split the input string into words.
	temp := strings.Split(input, " ")

	// Capitalize the first word.
	temp[0] = ToTitle(temp[0])

	// Join the words back to form the sentence and return.
	return strings.Join(temp, " ")
}
