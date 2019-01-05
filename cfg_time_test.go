package validation

import (
	"fmt"
	"time"
	"testing"
	//
	"github.com/jsonrouter/core/http"
)

func TestTime(t *testing.T) {

	req := http.NewMockRequest("GET", "/")

	b := true
	success := &b

	tests := map[string]*bool{
		"": nil,
		"01843 h88888": nil,
		"2017-06-23T00:00:00.000Z": success,
	}

	vc := Time(time.RFC3339Nano)

	for test, result := range tests {

		fmt.Println("TESTING:", test)

		if status, _ := vc.PathFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED: "+test)
			return
		}

	}

	for test, result := range tests {

		if status, _ := vc.BodyFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED: "+test)
			return
		}

	}
}
