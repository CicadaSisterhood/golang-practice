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
    // 1. Split the text into words
    text_split := strings.Split(text, " ")
    
    // 2. Count occurrences using a map
    occurrence_map := make(map[string]int)

    for _, s := range text_split{
        current_occurrences, exists := occurrence_map[s]
        if exists {
            occurrence_map[s] = current_occurrences + 1
        } else {
           occurrence_map[s] = 1 
        }
    }

    // 3. Return the map
    return occurrence_map
}
