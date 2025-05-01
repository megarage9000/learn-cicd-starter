package auth

import (
	"net/http"
	"testing"
)

func TestAuthHeader(t *testing.T) {

	type testCase struct {
		name     string
		input    http.Header
		expected string
	}

	normalHeader := http.Header{}
	normalHeader.Add("Authorization", "ApiKey 1234")

	emptyHeader := http.Header{}

	emptyKey := http.Header{}
	emptyKey.Add("Authorization", "")

	malformedHeader := http.Header{}
	malformedHeader.Add("Authorization", "2134214")

	tests := []testCase{
		{name: "Normal Header", input: normalHeader, expected: "1234"},
		{name: "No Header", input: emptyHeader, expected: ""},
		{name: "Empty Key", input: emptyKey, expected: ""},
		{name: "Malformed Header", input: malformedHeader, expected: ""},
	}

	for _, test := range tests {
		result, err := GetAPIKey(test.input)
		// If an error occured, but an expected was provided, fail
		if err != nil && test.expected != "" {
			t.Fatalf("%s Test Case Error: %v", test.name, err)
			// If the result does not equal expected, fail
		} else if test.expected != result {
			t.Fatalf("%s Test Case Expected: %v, Got: %v", test.name, test.expected, result)
		}
	}
}
