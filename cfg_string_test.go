package validation

import (
	"testing"
	//
	"github.com/jsonrouter/core/http"
)

func TestString(t *testing.T) {

	req := http.NewMockRequest("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"": nil,
		"a": success,
		"hello world": nil,
		"957!464&5": success,
		"hello world iweufghoqiuweh oqiwhe fpiqhw fihqwe": nil,
		"hello world iweufghoqiuweh oqiwhe fpiqhw fihqwehello world iweufghoqiuweh oqiwhe fpiqhw fihqwehello world iweufghoqiuweh oqiwhe fpiqhw fihqwehello world iweufghoqiuweh oqiwhe fpiqhw fihqwe": nil,
	}

	vc := String(1, 9)

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
