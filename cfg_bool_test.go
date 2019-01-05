package validation

import (
	"testing"
	//
	"github.com/jsonrouter/core/http"
)

func TestBool(t *testing.T) {

	req := http.NewMockRequest("GET", "/")

	b := true
	success := &b

	pathTests := map[string]*bool{
		"": nil,
		"true": success,
		"false": success,
	}

	bodyTests := map[interface{}]*bool{
		"": nil,
		true: success,
		false: success,
	}

	vc := Bool()

	for test, result := range pathTests {

		if status, _ := vc.PathFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED")
			return
		}

	}

	for test, result := range bodyTests {

		if status, _ := vc.BodyFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED")
			return
		}

	}
}
