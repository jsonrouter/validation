package validation

import (
	"fmt"
	"strings"
	//
	"github.com/jsonrouter/core/http"
)

// Returns a validation object that checks to see if it can resolve to a country struct
func CountryISO2() *Config {

	min := 2.0
	max := 2.0

	return NewConfig(
		"US",
		func (req http.Request, s string) (status *http.Status, _ interface{}) {

			status, s = checkString(
				req,
				min,
				max,
				strings.TrimSpace(strings.ToUpper(s)),
			)
			if status != nil {
				return status, nil
			}

			country := Countries()[s]

			if country == nil { status = req.Respond(400, "COUNTRY NOT FOUND: "+s) }

			return status, country
		},
		func (req http.Request, param interface{}) (status *http.Status, _ interface{}) {

			s, ok := param.(string)
			if !ok {
				status = req.Respond(
					400,
					fmt.Sprintf("COUNTRY PARAM NOT A STRING: %v", param),
				)
				return
			}

			status, s = checkString(
				req,
				min,
				max,
				strings.TrimSpace(strings.ToUpper(s)),
			)
			if status != nil {
				return
			}

			country := Countries()[s]
			if country == nil {
				status = req.Respond(400, "COUNTRY NOT FOUND: "+s)
				return
			}

			return nil, country
		},
		min,
		max,
	)
}
