package validation

import (
	"strings"
	//
	"github.com/jsonrouter/core/http"
)

// Username returns a validation object which checks for valid username
func Username(min, max float64) *Config {

	return NewConfig(
		"nameofsomething",
		func (req http.Request, s string) (status *http.Status, _ interface{}) {

			status, s = checkString(
				req,
				min,
				max,
				strings.TrimSpace(strings.ToLower(s)),
			)
			if status != nil {
				return status, nil
			}

			for _, char := range s {

				if !strings.Contains(USERNAME_CHARS, string(char)) { return req.Respond(400, ERR_INVALID_CHARS), nil }

			}

			return status, s
		},
		func (req http.Request, param interface{}) (status *http.Status, _ interface{}) {

			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			status, s = checkString(
				req,
				min,
				max,
				strings.TrimSpace(strings.ToLower(s)),
			)
			if status != nil {
				return status, nil
			}

			for _, char := range s {

				if !strings.Contains(USERNAME_CHARS, string(char)) { return req.Respond(400, ERR_INVALID_CHARS), nil }

			}

			return status, s
		},
		min,
		max,
	)
}
