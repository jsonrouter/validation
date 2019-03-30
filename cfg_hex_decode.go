package validation

import (
	"encoding/hex"
	//
	"github.com/jsonrouter/core/http"
)

// Returns a validation object that decodes a hex string with a length within optional range
func HexDecode(min, max float64) *Config {

	config := NewConfig(
		"0000000000000000000000000000000000000000000000000000000000000000",
		func (req http.Request, param string) (*http.Status, interface{}) {

			lp := float64(len(param))

			if lp < min || lp > max { return req.Respond(400, ERR_RANGE_EXCEED), nil }

			b, err := hex.DecodeString(param)
			if err != nil {
				return req.Respond(400, ERR_INVALID_CHARS), nil
			}

			return nil, b
		},
		func (req http.Request, param interface{}) (*http.Status, interface{}) {

			if min == 0 && param == nil { return nil, "" }

			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }

			lp := float64(len(s))

			if lp < min || lp > max { return req.Respond(400, ERR_RANGE_EXCEED), nil }

			b, err := hex.DecodeString(s)
			if err != nil {
				return req.Respond(400, ERR_INVALID_CHARS), nil
			}

			return nil, b
		},
	)

	config.Min = min
	config.Max = max

	return config
}
