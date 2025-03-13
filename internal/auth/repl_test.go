package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	tests := map[string]struct {
		input string
		sep   string
		want  []string
	}{
		"first":  {input: "ApiKey thisismykey", sep: " ", want: []string{"thisismykey"}},
		"second": {input: "ApiKey isthismykey", sep: " ", want: []string{"isthismykey"}},
		"third":  {input: "ApiKey whoknowsifthisismykey", sep: " ", want: []string{"whoknowsifthisismykey"}},
	}

	for name, testcase := range tests {
		header := make(http.Header)
		header.Add("Authorization", testcase.input)

		got, _ := GetAPIKey(header)

		if !reflect.DeepEqual(got, testcase.want[0]) {
			t.Errorf("Testcase: %v, Got %v, want %v", name, got, testcase.want[0])
		}
	}
}
