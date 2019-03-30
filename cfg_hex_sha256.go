package validation

import (
	"encoding/hex"
	//
	"github.com/jsonrouter/core/http"
)

// Returns a validation object that checks for a hex string with length of 64 chars.
func HexSHA256() *Config {

	max := float64(64)
	min := max

	config := NewConfig(
		"0000000000000000000000000000000000000000000000000000000000000000",
		func (req http.Request, param string) (*http.Status, interface{}) {

			lp := float64(len(param))

			if lp < min || lp > max { return req.Respond(400, ERR_RANGE_EXCEED), nil }

			if _, err := hex.DecodeString(param); err != nil {
				return req.Respond(400, ERR_INVALID_CHARS), nil
			}

			return nil, param
		},
		func (req http.Request, param interface{}) (*http.Status, interface{}) {

			if min == 0 && param == nil { return nil, "" }

			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			lp := float64(len(s))

			if lp < min || lp > max { return req.Respond(400, ERR_RANGE_EXCEED), nil }

			if _, err := hex.DecodeString(s); err != nil {
				return req.Respond(400, ERR_INVALID_CHARS), nil
			}

			return nil, s
		},
	)

	config.Min = min
	config.Max = max

	return config
}
