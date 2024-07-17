package url

import (
	"net/url"
	"strings"
)

// encode value param url
func QueryEscape(s string) string {
	return url.QueryEscape(strings.TrimSpace(s))
}
