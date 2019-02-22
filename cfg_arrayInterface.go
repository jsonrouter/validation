package validation

import (
	"github.com/jsonrouter/core/http"
)

// ArrayInterface returns a validation object for request body that checks a property to see if it's an array
func ArrayInterface() *Config {

	return NewConfig(
		[]interface{}{},
		nil,
		func (req http.Request, param interface{}) (*http.Status, interface{}) {

			m, ok := param.([]interface{}); if !ok { return req.Respond(400, ERR_NOT_ARRAY), nil }

			return nil, m
		},
	)
}
