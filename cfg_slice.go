package validation

import (
	//
	"github.com/jsonrouter/core/http"
)

// Returns a validation object that checks for a slice of integers
func Slice(vc *Config) *Config {

	return NewConfig(
		[]interface{}{},
		func (req http.Request, param string) (status *http.Status, _ interface{}) {

			panic("validation.Array(): THIS CODE SHOULD BE UNREACHABLE")

			return status, param
		},
		func (req http.Request, param interface{}) (status *http.Status, _ interface{}) {

			a, ok := param.([]interface{})
			if !ok {
				return req.Respond(400, ERR_NOT_ARRAY), nil
			}

			for x, item := range a {

				status, value := vc.BodyFunction(req, item)
				if status != nil {
					return status, nil
				}

				a[x] = value
			}

			return status, a
		},
	)
}
