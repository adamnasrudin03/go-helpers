package help

import (
	"math/rand"
	"strings"
	"time"
)

const (
	CharsetAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetAll      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()+,-.:;<=>?[]_{}"
)

// GenerateRandomString generates a random string of a specified length using the alphabet
// characters. It uses the current time as a seed for the random number generator.
//
// Parameters:
// - length: The length of the string to be generated.
//
// Returns:
// - A string of the specified length consisting of alphabet characters.
func GenerateRandomString(length int) string {
	// Define the charset for generating the random string
	const charset = CharsetAlphabet

	// Create a seeded random number generator using the current time as the seed
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create a byte slice to store the generated string
	b := make([]byte, length)

	// Generate the random string by selecting a random character from the charset for each index in the byte slice
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	// Return the generated string
	return string(b)
}

// GeneratePassword generates a random password of a specified length using a combination of
// lowercase letters, uppercase letters, numbers, and special characters.
//
// Parameters:
// - length: The length of the password to be generated.
//
// Returns:
// - A string of the specified length consisting of random characters.
func GeneratePassword(length int) string {
	// Define the charset for generating the password
	// The charset includes lowercase letters, uppercase letters, numbers, and special characters
	const charset = CharsetAll

	// Create a strings.Builder to store the generated password
	var password strings.Builder

	// Reserve memory for the password
	password.Grow(length)

	// Generate the password by selecting a random character from the charset for each index in the password
	for i := 0; i < length; i++ {
		// Select a random character from the charset and append it to the password
		password.WriteByte(charset[rand.Intn(len(charset))])
	}

	// Return the generated password
	return password.String()
}
