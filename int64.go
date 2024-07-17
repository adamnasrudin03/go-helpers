package help

// GetMinMaxInt64Array calculates the minimum and maximum values in an array of int64.
//
// Parameters:
// - params: An array of int64 values.
//
// Returns:
// - min: The minimum value in the array.
// - max: The maximum value in the array.
func GetMinMaxInt64Array(params []int64) (min int64, max int64) {
	// Initialize min and max with the first value in the array.
	if len(params) > 0 {
		min = params[0]
		max = params[0]
	}

	// Iterate over the array and update min and max if necessary.
	for _, value := range params {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	return min, max
}
