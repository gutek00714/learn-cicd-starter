package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError string
	}{
		{
			name:        "Valid ApiKey header",
			headers:     http.Header{"Authorization": []string{"ApiKey hello-world"}},
			expectedKey: "hello-world",
		},
		{
			name:          "Missing Authorization header",
			headers:       http.Header{},
			expectedError: "no authorization header included",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)

			if err != nil && err.Error() != tt.expectedError {
				t.Errorf("expected error %v, got %v", tt.expectedError, err.Error())
			}
			if err == nil && tt.expectedError != "" {
				t.Errorf("expected error %v, but got nil", tt.expectedError)
			}

			if key != tt.expectedKey {
				t.Errorf("expected key %v, got %v", tt.expectedKey, key)
			}
		})
	}
}
