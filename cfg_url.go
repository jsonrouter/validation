package validation

import (
	"strings"
	"net/url"
	//
	"github.com/jsonrouter/core/http"
)

// Url returns a validation object which checks for valid url
func URL() *Config {
	return NewConfig(
		"",
		func (req http.Request, param string) (*http.Status, interface{}) {
			param = strings.TrimSpace(param)
			_, err := url.ParseRequestURI(param); if err != nil { return req.Respond(400, ERR_PARSE_URL), "" }
			return nil, param
		},
		func (req http.Request, param interface{}) (*http.Status, interface{}) {
			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }
			s = strings.TrimSpace(strings.ToLower(s))
			_, err := url.ParseRequestURI(s); if err != nil { return req.Respond(400, ERR_PARSE_URL), "" }
			return nil, s
		},
	)
}
