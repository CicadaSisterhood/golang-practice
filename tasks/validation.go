package main

import (
	"fmt"
)

// return an error if index is out of bounds
func validateIndex(index int) error {
	if index < 0 || index >= len(tasks) {
		return fmt.Errorf("index %d is invalid", index)
	}
	return nil
}