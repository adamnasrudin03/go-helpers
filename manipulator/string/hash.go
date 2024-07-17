package string

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a hashed password from a plain text password using bcrypt.
// The function returns the hashed password and an error if any.
func HashPassword(password string) (hashed string, err error) {
	// Generate a hashed password from the plain text password using bcrypt.
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	// Convert the hashed password to a string and return it.
	return string(hash), nil
}

// PasswordValid checks if a given plain text password matches the hashed password.
// The function returns true if the passwords match, false otherwise.
func PasswordValid(hashPassword, password string) bool {
	// Convert the hashed password and plain text password to byte slices.
	hash, pass := []byte(hashPassword), []byte(password)

	// Compare the hashed password with the plain text password using bcrypt.
	// If the passwords match, return true, otherwise return false.
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}
