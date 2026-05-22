package main

import "testing"

func TestBuildPutResponse(t *testing.T) {
	testCases := []struct {
		name     string
		message  string
		expected string
	}{
		{
			name:     "count updated message",
			message:  "count updated",
			expected: `{"message":"count updated"}`,
		},
		{
			name:     "success message",
			message:  "success",
			expected: `{"message":"success"}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			result := buildPutResponse(testCase.message)

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
