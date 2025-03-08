package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAuthorization(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		want  string
	}{
		"simple": {
			input: http.Header{"Authorization": []string{"ApiKey 123456789"}},
			want:  "123456789",
		},
		"no key": {
			input: http.Header{"Authorization": []string{"ApiKey"}},
			want:  "fail",
		},
	}

	for name, tc := range tests {
		got, _ := GetAPIKey(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("%s: expected: %s got: %s", name, tc.want, got)
		}
	}
}
