package main

import (
	"fmt"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove query",
			inputURL: "https://blog.boot.dev/path?filter=name",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove trailing slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		// add more tests here
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
			fmt.Printf("test: '%s' passed\n", tc.name)
		})
	}
}

// func TestGetURLsFromHTML(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		inputURL  string
// 		inputBody string
// 		expected  []string
// 	}{
// 		{
// 			name:     "absolute and relative URLs",
// 			inputURL: "https://blog.boot.dev",
// 			inputBody: `
// 		<html>
// 			<body>
// 				<a href="path/one">
// 					<span>Boot.dev</span>
// 				</a>
// 				<a href="https://other.com/path/one">
// 					<span>Boot.dev</span>
// 				</a?
// 			</body>
// 		</html>
// 		`,
// 			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
// 		},
// 		{
// 			name:     "no href",
// 			inputURL: "https://blog.boot.dev",
// 			inputBody: `
// 		<html>
// 			<body>
// 				<a>
// 					<span>Boot.dev></span>
// 				</a>
// 			</body>
// 		</html>
// 		`,
// 			expected: nil,
// 		},
// 	}
// 	for i, tc := range tests {
// 		t.Run(tc.name, func(t *testing.T) {
// 			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
// 			if err != nil {
// 				t.Errorf("Test %v = '%s' FAIL: unexpected error: %v", i, tc.name, err)
// 				return
// 			}
// 			if !reflect.DeepEqual(actual, tc.expected) {
// 				t.Errorf("Test %v - %s FAIL: expected URLs: %v, actual: %v", i, tc.name, tc.expected, actual)
// 			}
// 			fmt.Printf("test : '%s' passed\n", tc.name)
// 		})
// 	}
// }
