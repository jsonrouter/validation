package validation

import (
	"html"
	"strings"
	//
	"github.com/microcosm-cc/bluemonday"
)

// Sanitize is used for string inputs to ensure some basic safety from injectioms or whatever (see bluemonday).
func Sanitize(s string) string {
	return html.UnescapeString(strings.TrimSpace(bluemonday.StrictPolicy().Sanitize(s)))
}
