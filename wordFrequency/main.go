package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

var wordRegexp = regexp.MustCompile(`\p{L}+\p{L}*`)

func main() {
	fmt.Println("Enter text:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	// Call your word frequency function
	freqMap, err := wordFrequency(text)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the results
	for word, count := range freqMap {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// wordFrequency takes a string and returns a map of word -> count
func wordFrequency(text string) (map[string]int, error) {
	// validate input
	err := validateInput(text)
	if err != nil {
		return nil, err
	}

	freq := make(map[string]int)

	// Find all word matches
	words := wordRegexp.FindAllString(text, -1)

	for _, word := range words {
		cleaned := strings.ToLower(word)
		if cleaned != "" {
			freq[cleaned]++
		}
	}

	return freq, nil
}

func validateInput(text string) error {
	if strings.TrimSpace(text) == "" {
		return fmt.Errorf("input string is empty")
	}

	const maxLength = 10000
	if len(text) > maxLength {
		return fmt.Errorf("input string exceeds %d characters", maxLength)
	}

	hasLetter := false
	for _, r := range text {
		if unicode.IsLetter(r) {
			hasLetter = true
			break
		}
	}
	if !hasLetter {
		return fmt.Errorf("input contains no letters")
	}

	return nil
}
