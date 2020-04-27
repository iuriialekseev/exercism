// Package isogram implements utilities to check if a string is an isogram
package isogram

import (
	"strings"
)

// IsIsogram returns true if given string is an isogram
func IsIsogram(str string) bool {
	charsOccurrences := make(map[rune]bool)

	for _, char := range strings.ToLower(str) {
		if char == '-' || char == ' ' {
			continue
		}

		if charsOccurrences[char] {
			return false
		}

		charsOccurrences[char] = true
	}

	return true
}
