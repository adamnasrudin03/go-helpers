package time

import "time"

// CheckTimeIsZeroToPointer returns a pointer to a time.Time if the input time is not zero,
// otherwise it returns nil.
//
// Parameters:
// - t: A time.Time object.
//
// Returns:
// - A pointer to a time.Time object.
func CheckTimeIsZeroToPointer(t time.Time) *time.Time {
	// Check if the input time is zero.
	if t.IsZero() {
		// If it is zero, return nil.
		return nil
	}

	// If it is not zero, return a pointer to the input time.
	return &t
}

// CheckTimePointerValue returns the value of a time.Time pointer.
// If the pointer is nil, it returns the zero value of time.Time.
//
// Parameters:
// - t: A pointer to a time.Time object.
//
// Returns:
// - The value of the time.Time object pointed to by the input pointer.
func CheckTimePointerValue(t *time.Time) time.Time {
	// Check if the input pointer is nil.
	if t == nil {
		// If it is nil, return the zero value of time.Time.
		return time.Time{}
	}

	// If the pointer is not nil, return the value of the time.Time object pointed to by the input pointer.
	return *t
}

// CheckTimeIsZeroToString converts a time.Time object to a string if it is not zero,
// otherwise it returns an empty string.
//
// Parameters:
// - t: A time.Time object.
// - formatDate: The desired format of the output string.
//
// Returns:
// - A string representation of the time.Time object if it is not zero, otherwise an empty string.
func CheckTimeIsZeroToString(t time.Time, formatDate string) string {
	// Check if the input time is zero.
	if t.IsZero() {
		// If it is zero, return an empty string.
		return ""
	}
	// If it is not zero, convert the time to the desired format and return it as a string.
	return t.Format(formatDate)
}

// CheckTimePointerToString converts a pointer to a time.Time object to a string,
// using the specified format. If the pointer is nil, it returns an empty string.
//
// Parameters:
// - t: A pointer to a time.Time object.
// - formatDate: The desired format of the output string.
//
// Returns:
// - A string representation of the time.Time object, formatted according to the specified format.
// - If the input pointer is nil, an empty string is returned.
func CheckTimePointerToString(t *time.Time, formatDate string) string {
	// Check if the input pointer is nil.
	if t == nil {
		// If it is nil, return an empty string.
		return ""
	}

	// If the pointer is not nil, convert the time to the desired format and return it as a string.
	return t.Format(formatDate)
}
