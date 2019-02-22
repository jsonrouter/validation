package validation

import (
	"strings"
	//
	"github.com/jsonrouter/core/http"
)

var validation_languages_map map[string]*Language

// LanguageISO2 returns a validation object that checks to see if it can resolve to a country struct
func LanguageISO2() *Config {

	if validation_languages_map == nil {

		validation_languages_map = Languages()
	}

	return NewConfig(
		"EN",
		func (req http.Request, param string) (status *http.Status, _ interface{}) {

			if len(param) != 2 { return req.Respond(400, "PARAM IS LONGER THAN 2 CHARS"), nil }

			param = strings.ToUpper(param)

			lang := validation_languages_map[param]

			if lang == nil { status = req.Respond(400, "LANGUAGE NOT FOUND: "+param) }

			return status, lang
		},
		func (req http.Request, param interface{}) (status *http.Status, _ interface{}) {

			s, ok := param.(string); if !ok { return nil, nil }

			if len(s) != 2 { return req.Respond(400, "PARAM IS LONGER THAN 2 CHARS"), nil }

			s = strings.ToUpper(s)

			lang := validation_languages_map[s]

			if lang == nil { status = req.Respond(400, "LANGUAGE NOT FOUND: "+s) }

			return status, lang
		},
	)
}
