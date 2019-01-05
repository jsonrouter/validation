package validation

import (
	"testing"
	//
	"github.com/jsonrouter/core/http"
)

func TestLanguageISO2(t *testing.T) {

	req := http.NewMockRequest("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"":	nil,
		"EN": success,
	}

	vc := LanguageISO2()

	for test, result := range tests {

		if status, _ := vc.BodyFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED")
		}

	}

}
