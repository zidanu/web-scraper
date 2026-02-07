package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://google.com/maps",
			expected: "google.com/maps",
		},
		{
			name:     "remove end slash",
			inputURL: "https://google.com/maps/",
			expected: "google.com/maps",
		},
		{
			name:     "remove http scheme",
			inputURL: "http://google.com/maps",
			expected: "google.com/maps",
		},
		{
			name:     "empty url",
			inputURL: "",
			expected: "",
		},
		{
			name:     "multiple end slashes",
			inputURL: "https://google.com/maps////",
			expected: "google.com/maps",
		},
		{
			name:     "ultimate test 1",
			inputURL: "https://blog.example.com:443/folder/page?type=article#comments",
			expected: "blog.example.com/folder/page?type=article",
		},
		{
			name:     "http default port stripped",
			inputURL: "http://google.com:80/maps/",
			expected: "google.com/maps",
		},
		{
			name:     "https non-default port kept",
			inputURL: "https://google.com:8443/maps/",
			expected: "google.com:8443/maps",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
