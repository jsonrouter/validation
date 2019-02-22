package validation

import (
	"encoding/base64"
	//
	"github.com/jsonrouter/core/http"
)

// Base64Decode returns a validation object which checks for valid username
func Base64Decode() *Config {

	return NewConfig(
		"",
		func (req http.Request, s string) (status *http.Status, _ interface{}) {

      		b, err := base64.StdEncoding.DecodeString(s); if err != nil { status = req.Respond(400, err.Error()) }

			return status, b
		},
		func (req http.Request, param interface{}) (status *http.Status, _ interface{}) {

			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			b, err := base64.StdEncoding.DecodeString(s); if err != nil { status = req.Respond(400, err.Error()) }

			return status, b
		},
	)
}
