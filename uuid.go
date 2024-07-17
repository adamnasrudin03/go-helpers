package help

import "github.com/google/uuid"

// GenerateUUID generates a new UUID v7
//
// It returns a UUID v7 and an error, if any.
func GenerateUUID() (uuid.UUID, error) {
	// Generate a new V7 (random) UUID
	// We use V7 for it's shorter length and it's use for
	// generating UUIDs from random data.
	generatedUUID, err := uuid.NewV7()
	if err != nil {
		// Return a zero value UUID and the error
		return uuid.UUID{}, err
	}
	// Return the generated UUID and nil error
	return generatedUUID, nil
}
