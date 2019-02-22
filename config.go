package validation

import 	(
	"regexp"
	"reflect"
	"strings"
	"encoding/hex"
	//
	"golang.org/x/crypto/sha3"
	//
	"github.com/jsonrouter/core/http"
)

// IsAlpha checkes whether a string is alphabetic.
func IsAlpha(s string) (bool, string) {
	return regexp.MustCompile("[a-zA-Z]").MatchString(s), strings.ToLower(s)
}

// IsAlphanumeric checks if a string is alphanumeric.
func IsAlphanumeric(s string) (bool, string) {
	return regexp.MustCompile("[a-zA-Z0-9_]").MatchString(s), strings.ToLower(s)
}

// Hash256 is a tool to exec the SHA25 on strings, taking a string input, and outputting hex.
func Hash256(input string) string {

	b := make([]byte, 64)

	sha3.ShakeSum256(b, []byte(input))

	return hex.EncodeToString(b)
}

type BodyValidationFunction func (http.Request, interface{}) (*http.Status, interface{})
type PathValidationFunction func (http.Request, string) (*http.Status, interface{})

type Spec struct {
	Type string
	Keys []string
}

type Config struct {
	Model interface{} `json:"model"`
	Type string `json:"type"`
	PathFunction PathValidationFunction `json:"-"`
	BodyFunction BodyValidationFunction `json:"-"`
	Keys []string `json:"-"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	RequiredValue bool `json:"required"`
	SummaryValue string `json:"summary"`
	DefaultValue interface{} `json:"default"`
	DescriptionValue string `json:"description"`
}

// Default adds a default value to the config
func (vc *Config) Default(x interface{}) *Config {
	if reflect.TypeOf(x).String() == vc.Type {
		vc.DefaultValue = x
	} else {
		panic("Wrong type '" + reflect.TypeOf(x).String() + "' for default value '" + vc.Type + "'")
	}
	return vc
}

// Description adds a description to the config
func (vc *Config) Description(x string) *Config {
	vc.DescriptionValue = x
	return vc
}

// Summary adds a description to the config
func (vc *Config) Summary(x string) *Config {
	vc.SummaryValue = x
	return vc
}

// Required signifies that the field is 'required'
func (vc *Config) Required() *Config {
	vc.RequiredValue = true
	return vc
}

// Key returns the first key associated with the path parameter
func (vc *Config) Key() string {
	return vc.Keys[0]
}

// KeyJoin joins the keys that are accocuated to a validation config, and concatonates them.
func (vc *Config) KeyJoin(delim string) string {

	return strings.Join(vc.Keys, delim)
}

// Expecting constructs a string that indicates what the required values were.
func (vc *Config) Expecting() string {

	return "expecting: " + vc.Type + " for keys: " + vc.KeyJoin(", ")
}

// NewConfig creates a new validation config with the supplied parameters.
func NewConfig(validationType interface{}, pathFunction PathValidationFunction, bodyFunction BodyValidationFunction, ranges ...float64) *Config {

	cfg := &Config{
		Model: validationType,
		Type: reflect.TypeOf(validationType).String(),
		PathFunction: pathFunction,
		BodyFunction: bodyFunction,
	}

	switch len(ranges) {

		case 2:
			cfg.Min = ranges[0]
			cfg.Max = ranges[1]

	}

	return cfg
}

type JSON struct {}

// Json returns a validation object which checks for (in)valid json
func Json() *Config {
	return NewConfig(
		JSON{},
		func (req http.Request, param string) (*http.Status, interface{}) {
			return nil, param
		},
		func (req http.Request, param interface{}) (*http.Status, interface{}) {
			s, ok := param.(string); if !ok { return req.Respond(400, ERR_NOT_STRING), nil }
			return nil, s
		},
	)
}
