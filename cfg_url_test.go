package validation

import (
	"testing"
	//
	"github.com/jsonrouter/core/http"
)

func TestURL(t *testing.T) {

	req := http.NewMockRequest("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"": nil,
		"http://ab.c": success,
		"ab": nil,
		"abc": nil,
	}

	vc := URL()

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
