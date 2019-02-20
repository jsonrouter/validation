package validation

import 	(
    "strings"
    "unicode"
	//
	"github.com/ttacon/libphonenumber"
	//
	"github.com/jsonrouter/core/http"
)

func validatePhoneNumber(req http.Request, countryCode, number string) (*http.Status, string) {

	number = strings.Replace(number, " ", "", -1)

	for _, s := range number {
		if string(s) != "+" && !unicode.IsNumber(s) { return req.Respond("PARAM CONTAINS INVALID CHARS FOR PHONE NUMBER"), "" }
	}

	if len(number) == 0 { return nil, "" }

	if len(number) > 16 { return req.Respond(400, "PARAM STRING TOO LONG TO BE A PHONE NUMBER"), "" }

	num, err := libphonenumber.Parse(number, countryCode); if req.Log().Error(err) { return req.Respond("PARAM IS AN INVALID PHONE NUMBER"), "" }

	format := libphonenumber.Format(num, libphonenumber.NATIONAL)

	return nil, format
}

// PhoneNumber returns a validation object that checks to see if a valid phone number is provided
func PhoneNumber(countryCode string) *Config {
		
	return NewConfig(
		"",
		func (req http.Request, number string) (*http.Status, interface{}) {

			return validatePhoneNumber(req, countryCode, number)
		},
		func (req http.Request, param interface{}) (*http.Status, interface{}) {

			number, ok := param.(string); if !ok { return req.Respond(400, "PARAM IS NOT A STRING"), nil }

			return validatePhoneNumber(req, countryCode, number)
		},
	)
}
