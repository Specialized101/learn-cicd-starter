package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "ApiKey 123456",
			expected: "123456",
		},
		{
			input:    "Bearer 123456",
			expected: "",
		},
		{
			input:    "ApiKey123456",
			expected: "",
		},
	}

	headers := http.Header{}
	for _, c := range cases {
		headers.Set("Authorization", c.input)
		actual, _ := GetAPIKey(headers)
		if actual != c.expected {
			t.Fatalf("expected: %v; got: %v\n", c.expected, actual)
		}
	}
}
