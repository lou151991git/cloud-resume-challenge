package main

import "testing"

func TestBuildResponse(t *testing.T) {
	testCases := []struct {
		name     string
		count    string
		expected string
	}{
		{
			name:     "count is 1",
			count:    "1",
			expected: `{"count":1}`,
		},
		{
			name:     "count is 999",
			count:    "999",
			expected: `{"count":999}`,
		},
		{
			name:     "count is 0",
			count:    "0",
			expected: `{"count":0}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			result := buildResponse(testCase.count)

			if result != testCase.expected {
				t.Errorf(
					"Expected %v, but got %v",
					testCase.expected,
					result,
				)
			}
		})
	}
}
