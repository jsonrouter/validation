package validation

import (
	//
	"github.com/jsonrouter/core/http"
)

// Returns a validation object that allows any value
func ArrayString() *Config {

	return NewConfig(
		[]string{},
		func (req http.Request, param string) (status *http.Status, _ interface{}) {

			// not really a possible scenario

			return status, param
		},
		func (req http.Request, param interface{}) (status *http.Status, _ interface{}) {

			a, ok := param.([]interface{})
			if !ok {
				return req.Respond(400, ERR_NOT_ARRAY), nil
			}

			b := make([]string, len(a))
			for x, item := range a {
				b[x], ok = item.(string)
				if !ok {
					return req.Respond(400, ERR_NOT_STRING), nil
				}
			}

			return status, b
		},
	)
}
