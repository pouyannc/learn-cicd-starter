package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header1 := http.Header{}
	header1.Add("Authorization", "ApiKey 12345")

	header2 := http.Header{}
	header2.Add("Authorization", "")

	header3 := http.Header{}

	tests := map[string]struct {
		headers    http.Header
		wantAPIKey string
		wantErr    bool
	}{
		"simple get API key":     {headers: header1, wantAPIKey: "12345", wantErr: false},
		"no authorization value": {headers: header2, wantAPIKey: "", wantErr: true},
		"no authorization key":   {headers: header3, wantAPIKey: "", wantErr: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotAPIKey, err := GetAPIKey(tc.headers)
			if tc.wantErr != (err != nil) {
				t.Fatalf("expected err: %v, got err: %v", tc.wantErr, err)
			}
			if gotAPIKey != tc.wantAPIKey {
				t.Fatalf("expected: %v, got: %v", tc.wantAPIKey, gotAPIKey)
			}
		})
	}
}
