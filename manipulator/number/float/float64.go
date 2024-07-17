package float

import "math"

// RoundUpFloat rounds up the given float64 to the given uint precision.
//
// Rounds up the given float64 to the given uint precision. For example, if
// the input is 12.345 and the precision is 2, this function will return
// 12.35.
//
// Parameters:
// - input: The float64 to be rounded up.
// - precision: The number of decimal places to round to.
//
// Returns:
// - The rounded up float64.
//
// check https://adamnasrudin.vercel.app/cheat-sheet/rounding-float-using-golang
func RoundUpFloat(input float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(input*ratio) / ratio
}

// RoundDownFloat rounds down the given float64 to the given uint precision.
//
// Rounds down the given float64 to the given uint precision. For example, if
// the input is 12.345 and the precision is 2, this function will return
// 12.34.
//
// Parameters:
// - input: The float64 to be rounded down.
// - precision: The number of decimal places to round to.
//
// Returns:
// - The rounded down float64.
//
// check https://adamnasrudin.vercel.app/cheat-sheet/rounding-float-using-golang
func RoundDownFloat(input float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Floor(input*ratio) / ratio
}

// RoundFloat rounds the given float64 to the given uint precision, based on
// the rounding mode.
//
// If roundingUp is true, this function will round the given float64 up.
// If roundingUp is false, this function will round the given float64 down.
//
// Parameters:
// - input: The float64 to be rounded.
// - precision: The number of decimal places to round to.
// - roundingUp: A boolean indicating whether to round up or down.
//
// Returns:
// - The rounded float64.
//
// Example:
// RoundFloat(12.345, 2, true) // 12.35
// RoundFloat(12.345, 2, false) // 12.34

//
// check https://adamnasrudin.vercel.app/cheat-sheet/rounding-float-using-golang
func RoundFloat(input float64, precision uint, roundingUp bool) float64 {
	if roundingUp {
		return RoundUpFloat(input, precision)
	} else {
		return RoundDownFloat(input, precision)
	}
}
