package validation

import (
	//
	"github.com/jsonrouter/core/http"
)

// ArrayFloat64 returns a validation object that checks for a slice of float64
func ArrayFloat64() *Config {

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

			b := make([]float64, len(a))
			for x, item := range a {
				i, ok := item.(float64)
				if !ok {
					return req.Respond(400, ERR_NOT_FLOAT64), nil
				}
				b[x] = i
			}

			return status, b
		},
	)
}
