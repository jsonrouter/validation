package validation

import (
	"strings"
	//
	"github.com/jsonrouter/core/http"
)

// Returns a validation object which checks for delimiter-separated string like CSV
func StringSplit(delimiter string) *Config {

	min := 0.0
	max := STRING_MAX_LENGTH

	return NewConfig(
		[]string{"a","b"},
		func (req http.Request, s string) (status *http.Status, _ interface{}) {

			status, s = checkString(req, min, max, s)
			if status != nil {
				return status, nil
			}

			list := []string{}

			for _, part := range strings.Split(s, delimiter) {

				if len(part) == 0 { continue }

				list = append(list, strings.TrimSpace(part))

			}

			return nil, list
		},
		func (req http.Request, param interface{}) (status *http.Status, _ interface{}) {

			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			status, s = checkString(req, min, max, s)
			if status != nil {
				return status, nil
			}

			list := []string{}

			for _, part := range strings.Split(s, delimiter) {

				if len(part) == 0 { continue }

				list = append(list, strings.TrimSpace(part))

			}

			return nil, list
		},
		min,
		max,
	)
}
