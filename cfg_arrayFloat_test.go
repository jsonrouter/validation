package validation

import (
	"testing"
	"encoding/json"
	//
	"github.com/jsonrouter/core/http"
)

func TestArrayFloat(t *testing.T) {

	req := http.NewMockRequest("GET", "/")

	b := true
	success := &b

	bodyValues := []interface{}{
		[]interface{}{"a", "b", "c"},
		[]interface{}{true, false, true},
		[]interface{}{3, true, 3.3},
		[]interface{}{3, true, 3},
		[]interface{}{10.3, 2.03, 3.3},
		[]interface{}{3, 2, 1},
		[]interface{}{3},
		[]interface{}{0, 0, 0, 0.1},
	}

	var a []interface{}
	if b, err := json.Marshal(bodyValues); err != nil {
		t.Error(err)
		t.Fail()
	} else {
		if err := json.Unmarshal(b, &a); err != nil {
			t.Error(err)
			t.Fail()
		}
	}

	bodyValues = a

	bodyResults := []*bool{
		nil,
		nil,
		nil,
		nil,
		success,
		success,
		success,
		success,
	}

	vc := ArrayFloat64()

	for i, result := range bodyResults {

		test := bodyValues[i].([]interface{})

		if status, _ := vc.BodyFunction(req, test); (result == nil && status == nil) || (result != nil && status != nil) {

			t.Error("FAILED:", test)
			t.Fail()
			break
		}

	}

}
