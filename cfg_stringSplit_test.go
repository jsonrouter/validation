package validation

import (
	"testing"
	//
	"github.com/jsonrouter/core/http"
)

func TestStringSplit(t *testing.T) {

	req := http.NewMockRequest("GET", "/")

	tests := map[string][]string{
		" yes , no , maybe ": []string{"yes", "no", "maybe"},
		"raining": []string{"raining"},
	}

	vc := StringSplit(",")

	for test, result := range tests {

		status, output := vc.PathFunction(req, test); if status != nil {
			t.Error("FAILED")
			return
		}

		if len(result) != len(output.([]string)) {
			t.Error("FAILED")
			return
		}

		for x, v := range output.([]string) {

			if result[x] != v {
				t.Error("FAILED")
				return
			}

		}
	}

	for test, result := range tests {

		status, output := vc.BodyFunction(req, test); if status != nil {
			t.Error("FAILED")
			return
		}

		if len(result) != len(output.([]string)) {
			t.Error("FAILED")
			return
		}

		for x, v := range output.([]string) {

			if result[x] != v {
				t.Error("FAILED")
				return
			}

		}
	}
}
