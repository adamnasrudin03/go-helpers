package help

// CheckArrayFloat64Nil checks if the input array of float64 is nil or empty.
// If not, it returns the input array.
// If the input array is nil or empty, it returns an empty array.
//
// Parameters:
// - input: An array of float64.
//
// Returns:
// - An array of float64.
func CheckArrayFloat64Nil(input []float64) []float64 {
	// Check if the length of the input array is greater than 0.
	if len(input) > 0 {
		// If the array is not empty, return the input array.
		return input
	}
	// If the array is empty or nil, return an empty array.
	return []float64{}
}

// CheckFloat64Value checks if the input pointer to a float64 is nil or empty.
// If not, it returns the value of the input pointer.
// If the input pointer is nil, it returns 0.0.
//
// Parameters:
// - input: A pointer to a float64.
//
// Returns:
// - A float64.
func CheckFloat64Value(input *float64) float64 {
	// Check if the input pointer is nil.
	if input == nil {
		// If the input pointer is nil, return 0.0.
		return 0.0
	}
	// If the input pointer is not nil, return the value of the input pointer.
	return *input
}
