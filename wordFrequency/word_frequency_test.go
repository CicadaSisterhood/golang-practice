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
                "go": 4,
                "is": 1,
                "fun": 1,
            },
        },
        {
            input: "Wait—what? Yes, wait!",
            expected: map[string]int{
                "wait": 2,
                "what": 1,
                "yes": 1,
            },
        },
    }

    for _, test := range tests {
        got := wordFrequency(test.input)
        if !reflect.DeepEqual(got, test.expected) {
            t.Errorf("wordFrequency(%q) = %v; want %v", test.input, got, test.expected)
        }
    }
}
