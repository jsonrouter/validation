package validation

import (
	//
	"github.com/jsonrouter/core/http"
)

// Bool returns a validation object that checks for a bool which parses correctly
func Bool() *Config {

	return NewConfig(
		false,
		func (req http.Request, param string) (*http.Status, interface{}) {

			switch param {

				case "true":	return nil, true
				case "false":	return nil, false

			}

			return req.Respond(400, ERR_NOT_BOOL), nil
		},
		func (req http.Request, param interface{}) (*http.Status, interface{}) {

			b, ok := param.(bool); if !ok { return req.Respond(400, ERR_NOT_BOOL), nil }

			return nil, b
		},
	)
}