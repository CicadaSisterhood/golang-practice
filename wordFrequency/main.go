package main

import (
	"fmt"
	"strings"
)

func main() {
    // Hardcoded input string
    text := "the cat sat on the mat"

    // Call your word frequency function
    freqMap := wordFrequency(text)

    // Print the results
    for word, count := range freqMap {
        fmt.Printf("%s: %d\n", word, count)
    }
}

// wordFrequency takes a string and returns a map of word -> count
func wordFrequency(text string) map[string]int {
    text = strings.ToLower(text)
    words := strings.Fields(text)
    freq := make(map[string]int)

    for _, word := range words {
        freq[word]++
    }

    return freq
}
