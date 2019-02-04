package validation

import (
	"fmt"
	//
	"github.com/jsonrouter/core/http"
)

func checkString(req http.Request, min, max float64, s string) (*http.Status, string) {

	s = Sanitize(s)

	length := len(s)

	if length < int(min) {
		return req.Respond(
			400,
			fmt.Sprintf("%s: min (%d)", ERR_RANGE_EXCEED, length),
		), ""
	}

	if length > int(max) {
		return req.Respond(
			400,
			fmt.Sprintf("%s: max (%d)", ERR_RANGE_EXCEED, length),
		), ""
	}

	return nil, s
}

// Returns a validation object that checks for a string with a length within optional range
func String(min, max float64) *Config {

	return NewConfig(
		"",
		func (req http.Request, param string) (*http.Status, interface{}) {

			return checkString(req, min, max, param)
		},
		func (req http.Request, param interface{}) (*http.Status, interface{}) {

			if min == 0 && param == nil { return nil, "" }

			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			return checkString(req, min, max, s)
		},
		min,
		max,
	)
}
