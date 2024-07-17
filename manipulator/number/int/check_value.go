package int

// CheckArrayIntNil checks if the input array of integers is nil or empty. If not, it returns the input array.
// If the input array is nil or empty, it returns an empty array.
//
// Parameters:
// - input: An array of integers.
//
// Returns:
// - An array of integers.
func CheckArrayIntNil(input []int) []int {
	// Check if the length of the input array is greater than 0.
	if len(input) > 0 {
		// If the array is not empty, return the input array.
		return input
	}
	// If the array is empty or nil, return an empty array.
	return []int{}
}

// CheckIntValue checks if the input pointer of integer is nil or empty. If not, it returns the value of the input pointer.
// If the input pointer is nil or empty, it returns 0.
//
// Parameters:
// - input: A pointer of integer.
//
// Returns:
// - An integer.
func CheckIntValue(input *int) int {
	// Check if the input pointer is nil.
	if input == nil {
		// If the input pointer is nil, return 0.
		return 0
	}
	// If the input pointer is not nil, return the value of the input pointer.
	return *input
}
