package validation

import (
	"testing"
	//
	"github.com/jsonrouter/core/http"
)

func TestHexSHA1(t *testing.T) {

	req := http.NewMockRequest("GET", "/")

	b := true
	success := &b

	pathTests := map[string]*bool{
		"000000000000000000000000000000000000000000": nil,
		"00000000000000000000000000000000000000000": nil,
		"0000000000000000000000000000000000000000": success,
		"000000000000000000000000000000000000000": nil,
		"00000000000000000000000000000000000000": nil,
		"": nil,
		"00": nil,
		"hello world": nil,
		"0.1": nil,
		"67957!464&5": nil,
		"2": nil,
		"true": nil,
	}

	vc := HexSHA1()

	for test, result := range pathTests {

		if status, _ := vc.PathFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED")
			t.Fail()
		}

	}

	for test, result := range pathTests {

		if status, _ := vc.BodyFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED")
			t.Fail()
		}

	}
}
