package validation

import (
	"testing"
	//
	"github.com/jsonrouter/core/http"
)

func TestIPv4Address(t *testing.T) {

	req := http.NewMockRequest("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"": nil,
		"a": nil,
		"1...": nil,
		"1.2": nil,
		"1.2.3": nil,
		"1.2.3.4": success,
		"255.255.255.255": success,
	}

	vc := IPv4Address()

	for test, result := range tests {

		if status, _ := vc.PathFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED:", test)
			return
		}

	}

	for test, result := range tests {

		if status, _ := vc.BodyFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED:", test)
			return
		}

	}
}
