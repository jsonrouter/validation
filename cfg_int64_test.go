package validation

import (
	"fmt"
	"testing"
	//
	"github.com/jsonrouter/core/http"
)

func TestInt64(t *testing.T) {

	req := http.NewMockRequest("GET", "/")

	b := true
	success := &b

	pathTests := map[string]*bool{
		"0": success,
		"17808779797": success,
		"hello world": nil,
		"0.1": nil,
		"67957!464&5": nil,
	}

	bodyTests := map[interface{}]*bool{
		0: success,
		int64(178087797977587): success,
		"hello world": nil,
		0.1: success,
		"67957!464&5": nil,
	}

	vc := Int64()

	for test, result := range pathTests {

		fmt.Println("PATH TESTING:", test)

		if status, _ := vc.PathFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED")
		}

	}

	for test, result := range bodyTests {

		fmt.Println("BODY TESTING:", test)

		if status, _ := vc.BodyFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			fmt.Println("FAILED", result, status)

			t.Error("FAILED")
		}

		fmt.Println("PASSED", test)

	}
}
