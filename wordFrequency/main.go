package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var wordRegexp = regexp.MustCompile(`\p{L}+\p{L}*`)


func main() {
    fmt.Println("Enter text:")

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()

    // Call your word frequency function
    freqMap := wordFrequency(text)

    // Print the results
    for word, count := range freqMap {
        fmt.Printf("%s: %d\n", word, count)
    }
}

// wordFrequency takes a string and returns a map of word -> count
func wordFrequency(text string) map[string]int {
    freq := make(map[string]int)
    
    // Find all word matches
    words := wordRegexp.FindAllString(text, -1)

    for _, word := range words {
        cleaned := strings.ToLower(word)
        if cleaned != ""{
            freq[cleaned]++
        }
    }

    return freq
}
