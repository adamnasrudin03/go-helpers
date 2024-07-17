package int

import (
	"math"
	"math/rand"
)

// GenerateRandomNumber generates a random number within a specified length.
// The length parameter determines the maximum value of the generated number.
//
// Parameters:
// - length: The length of the number range.
//
// Returns:
// - A randomly generated number within the specified range.
// see https://adamnasrudin.vercel.app/cheat-sheet/generate-random-number-using-golang
func GenerateRandomNumber(length int) int {
	// Generate a random number within the range of 10^length
	// using the math/rand package.

	// Calculate the maximum value of the range based on the length.
	maxValue := int(math.Pow10(length))

	// Generate a random number within the range using rand.Intn.
	num := rand.Intn(maxValue)

	// Return the generated random number.
	return num
}
