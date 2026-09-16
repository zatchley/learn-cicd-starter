package auth

import (
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input map[string][]string
		want  string
	}{
		"simple":           {input: map[string][]string{"Authorization": []string{"ApiKey key"}}, want: "key"},
		"no auth header":   {input: map[string][]string{}, want: ""},
		"bad auth header1": {input: map[string][]string{"Authorization": []string{"ApiKey"}}, want: ""},
		"bad auth header2": {input: map[string][]string{"Authorization": []string{"BadKey key"}}, want: ""},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := GetAPIKey(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
			}
		})
	}
}
