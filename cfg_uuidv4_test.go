package validation

import (
	"testing"
	//
	"github.com/jsonrouter/core/http"
)

func TestUUIDv4(t *testing.T) {

	req := http.NewMockRequest("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"62c2639c-7b60-4403-81f0-eb40a01ae1bb": success,
		"98ceed88-756e-4823-96ef-2815eafc0c1e": success,
		"98ceed88-756e-4823-96ef-2815eafc0c1j": nil,
		"98ceed88-756e-4823-96ef-2815eafc0c1": nil,
		"98ceed88-756e-4823-96ef2815eafc0c1e": nil,
		"a": nil,
	}

	vc := UUIDv4()

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
