package validation

import (
	"testing"
	//
	"github.com/jsonrouter/core/http"
)

func TestCountryISO3(t *testing.T) {

	req := http.NewMockRequest("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"":	nil,
		"GBR": success,
		"DNK": success,
		"ESP": success,
		"EEE": nil,
		"USAF": nil,
		"USA": success,
		"US": nil,
		"GB": nil,
	}

	vc := CountryISO3()

	for test, result := range tests {

		t.Run(
			"TESTING alpha3 COUNTRY: "+test,
			func (t *testing.T) {

				status, value := vc.BodyFunction(req, test)
				if status == nil {
					if value == nil {
						if result == nil {
							t.Log("(FAILED AS REQUIRED)")
							return
						}
						t.Errorf("INVALID COUNTRY AND STATUS VALUE: %v %v", value, status)
						t.Fail()
					}
					c, ok := value.(*Country)
					if ok {
						if test != c.Alpha3Code {
							if result == nil {
								t.Log("(FAILED AS REQUIRED)")
							} else {
								t.Errorf("INVALID COUNTRY VALUE: %v %v", value, status)
								t.Fail()
							}
						}
					}
					return
				}
				if result == nil {
					t.Log("(FAILED AS REQUIRED)")
					return
				}
				t.Errorf("INVALID STATUS VALUE: %v", status)
				t.Fail()
				return
			},
		)

	}

	for test, result := range tests {

		t.Run(
			"TESTING alpha3 COUNTRY: "+test,
			func (t *testing.T) {

				status, value := vc.BodyFunction(req, test)
				if status == nil {
					if value == nil {
						if result == nil {
							t.Log("(FAILED AS REQUIRED)")
							return
						}
						t.Errorf("INVALID COUNTRY AND STATUS VALUE: %v %v", value, status)
						t.Fail()
					}
					c, ok := value.(*Country)
					if ok {
						if test != c.Alpha3Code {
							if result == nil {
								t.Log("(FAILED AS REQUIRED)")
							} else {
								t.Errorf("INVALID COUNTRY VALUE: %v %v", value, status)
								t.Fail()
							}
						}
					}
					return
				}
				if result == nil {
					t.Log("(FAILED AS REQUIRED)")
					return
				}
				t.Errorf("INVALID STATUS VALUE: %v", status)
				t.Fail()
				return
			},
		)

	}
}
