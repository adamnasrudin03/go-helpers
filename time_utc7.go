package help

import "time"

// OptionTimeUTC+7 is a function type used for applying options to TimeUTC7 instances.
type OptionTimeUTC7 func(*TimeUTC7)

// TimeUTC7 represents a time with its location set to UTC+7, specifically for Western Indonesia Time.
type TimeUTC7 struct {
	loc *time.Location // loc holds the location set to UTC+7.
}

// NewTimeUTC7 creates and returns a new TimeUTC7 instance.
// It initializes the location to Jakarta (UTC+7) by default.
// Additional options can be passed to customize the TimeUTC7 instance.
//
// Returns:
// - A pointer to the newly created TimeUTC7 instance.
func NewTimeUTC7(options ...OptionTimeUTC7) *TimeUTC7 {
	loc, _ := time.LoadLocation(AsiaJakarta) // Load the Jakarta timezone for UTC+7.
	t := &TimeUTC7{
		loc: loc,
	}
	// Apply any passed options to the TimeUTC7 instance.
	for _, o := range options {
		o(t)
	}
	return t
}

// Now returns the current time adjusted to the TimeUTC7's location (UTC+7).
//
// Returns:
// - A time.Time value representing the current time in UTC+7.
func (t *TimeUTC7) Now() time.Time {
	return time.Now().In(t.loc) // Adjust the current time to UTC+7.
}

// ParseUTC7 attempts to parse a given string into a time.Time value based on the specified format.
// The parsing is done considering the UTC+7 timezone.
// If parsing fails, the current time in UTC+7 and an error are returned.
//
// Parameters:
// - timeFormat: The format for parsing the input string.
// - value: The string value to parse into a time.Time.
//
// Returns:
// - A time.Time value representing the parsed time in UTC+7.
// - An error if the parsing fails.
func (t *TimeUTC7) ParseUTC7(timeFormat string, value string) (time.Time, error) {
	timeUTC7, err := time.ParseInLocation(timeFormat, value, t.loc)
	if err != nil {
		return t.Now(), err // Return current time in UTC+7 and the error if parsing fails.
	}

	return timeUTC7, nil
}

// StartDate computes the start of the day (00:00:00) for a given time in UTC+7.
//
// Parameters:
// - input: The time for which to find the start of the day.
//
// Returns:
// - A time.Time value representing the start of the day in UTC+7.
func (t *TimeUTC7) StartDate(input time.Time) time.Time {
	startOfDay, _ := t.ParseUTC7(FormatDate, input.Format(FormatDate)) // Parse the input time to get the start of the day.
	return startOfDay
}

// EndDate computes the end of the day (23:59:59) for a given time in UTC+7.
//
// Parameters:
// - input: The time for which to find the end of the day.
//
// Returns:
// - A time.Time value representing the end of the day in UTC+7.
func (t *TimeUTC7) EndDate(input time.Time) time.Time {
	endOfDay := t.StartDate(input).Add(23*time.Hour + 59*time.Minute + 59*time.Second) // Calculate 23:59:59 of the input day.
	return endOfDay
}

// StartDateInString parses a given date string to find the start of the day (00:00:00) in UTC+7.
//
// Parameters:
// - inputDate: The date string to parse.
//
// Returns:
// - A time.Time value representing the start of the day for the parsed date in UTC+7.
func (t *TimeUTC7) StartDateInString(inputDate string) time.Time {
	startOfDay, _ := t.ParseUTC7(FormatDate, inputDate) // Parse the date string to get the start of the day.
	return startOfDay
}

// EndDateInString parses a given date string to find the end of the day (23:59:59) in UTC+7.
//
// Parameters:
// - inputDate: The date string to parse.
//
// Returns:
// - A time.Time value representing the end of the day for the parsed date in UTC+7.
func (t *TimeUTC7) EndDateInString(inputDate string) time.Time {
	endOfDay := t.StartDateInString(inputDate).Add(23*time.Hour + 59*time.Minute + 59*time.Second) // Calculate 23:59:59 of the parsed date.
	return endOfDay
}
