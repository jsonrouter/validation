package validation

import (
	"strconv"
	//
	"github.com/jsonrouter/core/http"
)

// Int64 returns a validation object that checks for an int64 which parses correctly
func Int64() *Config {

	return NewConfig(
		int64(1),
		func (req http.Request, param string) (status *http.Status, _ interface{}) {

			val, err := strconv.ParseInt(param, 10, 64); if req.Log().Error(err) { status = req.Respond(400, ERR_PARSE_INT64) }

			return status, val
		},
		func (req http.Request, param interface{}) (status *http.Status, _ interface{}) {

			f, ok := param.(float64)

			if !ok {

				ii, ok := param.(int64)

				if !ok {

					i, ok := param.(int)

					if !ok { status = req.Respond(400, ERR_PARSE_INT64) }

					return status, i
				}

				return status, ii
			}

			return status, int64(f)
		},
	)
}
