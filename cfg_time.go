package validation

import (
	"time"
	//
	"github.com/jsonrouter/core/http"
)

// SQLTimestamp returns a validation object which checks for valid time
func SQLTimestamp() *Config {

	return NewConfig(
		time.Now(),
		func (req http.Request, param string) (*http.Status, interface{}) {

			return nil, param
		},
		func (req http.Request, param interface{}) (*http.Status, interface{}) {

			return nil, param
		},
	)
}

// Now returns a validation object which outputs the current time
func Now() *Config {

	return NewConfig(
		time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC).UnixNano(),
		func (req http.Request, param string) (*http.Status, interface{}) {

			return nil, time.Now().UTC().UnixNano()
		},
		func (req http.Request, param interface{}) (*http.Status, interface{}) {

			return nil, time.Now().UTC().UnixNano()
		},
	)
}

// Time returns a validation object which checks for valid time like 2017-06-23T00:00:00.000Z
func Time(layout string) *Config {

	return NewConfig(
		"",
		func (req http.Request, param string) (status *http.Status, _ interface{}) {

			return verifyTime(req, layout, param)
		},
		func (req http.Request, param interface{}) (status *http.Status, _ interface{}) {

			s, ok := param.(string); if !ok { status = req.Respond(400, ERR_NOT_STRING) }

			return verifyTime(req, layout, s)
		},
	)
}

func verifyTime(req http.Request, layout, input string) (*http.Status, string) {

    _, err := time.Parse(layout, input); if req.Log().Error(err) { return req.Respond(400, ERR_PARSE_TIME), "" }

    return nil, input
}