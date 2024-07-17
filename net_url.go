package help

import (
	"net/url"
	"strings"
)

// QueryEscape is a wrapper around the net/url.QueryEscape function.
// It trims the input string and then encodes it using the URL encoding.
//
// s: the string to be encoded
// returns: the encoded string
func QueryEscape(s string) string {
	// Trim the input string to remove any leading or trailing white spaces.
	trimmed := strings.TrimSpace(s)

	// Encode the trimmed string using the URL encoding.
	// The QueryEscape function from the net/url package is used for this.
	encoded := url.QueryEscape(trimmed)

	return encoded
}
