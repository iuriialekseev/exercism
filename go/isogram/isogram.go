// Package isogram implements utilities to check if a string is an isogram
package isogram

import (
	"strings"
)

// IsIsogram returns true if given string is an isogram
func IsIsogram(str string) bool {
	charsOccurrences := make(map[rune]bool)

	for _, char := range strings.ToLower(str) {
		if _, ok := charsOccurrences[char]; ok {
			return false
		}

		if char != '-' && char != ' ' {
			charsOccurrences[char] = true
		}
	}

	return true
}
