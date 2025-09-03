package main

import (
	"reflect"
	"testing"
)

func TestWordFrequencyWithPunctuation(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input: "Hello, world! Hello...",
			expected: map[string]int{
				"hello": 2,
				"world": 1,
			},
		},
		{
			input: "Go is fun. Go, go, go!",
			expected: map[string]int{
				"go":  4,
				"is":  1,
				"fun": 1,
			},
		},
		{
			input: "Wait—what? Yes, wait!",
			expected: map[string]int{
				"wait": 2,
				"what": 1,
				"yes":  1,
			},
		},
	}

	for _, test := range tests {
		got, _ := wordFrequency(test.input)
		if !reflect.DeepEqual(got, test.expected) {
			t.Errorf("wordFrequency(%q) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestWordFrequencyErrors(t *testing.T) {
	testCases := []struct {
		input       string
		description string
	}{
		{"", "empty string"},
		{"   ", "whitespace only"},
		{"1234---", "no letters, only numbers/punctuation"},
	}

	for _, tc := range testCases {
		_, err := wordFrequency(tc.input)
		if err == nil {
			t.Errorf("Expected error for %s, got nil", tc.description)
		}
	}
}
