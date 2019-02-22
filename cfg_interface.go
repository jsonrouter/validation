package validation

import (
	//
	"github.com/jsonrouter/core/http"
)

// Interface returns a validation object that allows any value
func Interface() *Config {

	return NewConfig(
		0,
		func (req http.Request, param string) (status *http.Status, _ interface{}) {

			return status, param
		},
		func (req http.Request, param interface{}) (status *http.Status, _ interface{}) {

			return status, param
		},
	)
}
