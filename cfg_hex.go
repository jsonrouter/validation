package validation

import (
	"encoding/hex"
	//
	"github.com/jsonrouter/core/http"
)

// Hex returns a validation object that checks for a hex string with a length within optional range
func Hex(min, max float64) *Config {

	config := NewConfig(
		"0000000000000000000000000000000000000000000000000000000000000000",
		func (req http.Request, param string) (*http.Status, interface{}) {

			lp := float64(len(param))

			if lp < min || lp > max { return req.Respond(400, ERR_RANGE_EXCEED), nil }

			_, err := hex.DecodeString(param)
			if err != nil {
				return req.Respond(400, ERR_INVALID_CHARS), nil
			}

			return nil, param
		},
		func (req http.Request, param interface{}) (*http.Status, interface{}) {

			if min == 0 && param == nil { return nil, "" }

			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			lp := float64(len(s))

			if lp < min || lp > max { return req.Respond(400, ERR_RANGE_EXCEED), nil }

			_, err := hex.DecodeString(s)
			if err != nil {
				return req.Respond(400, ERR_INVALID_CHARS), nil
			}

			return nil, s
		},
	)

	config.Min = min
	config.Max = max

	return config
}
