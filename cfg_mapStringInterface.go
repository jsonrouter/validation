package validation

import (
	"github.com/jsonrouter/core/http"
)

// MapStringInterface returns a validation object for request body that checks a property to see if it's an object
func MapStringInterface() *Config {

	return NewConfig(
		map[string]interface{}{},
		nil,
		func (req http.Request, param interface{}) (*http.Status, interface{}) {

			x, ok := param.(map[string]interface{}); if !ok { return req.Respond(400, ERR_NOT_OBJECT), nil }

			return nil, x
		},
	)
}
