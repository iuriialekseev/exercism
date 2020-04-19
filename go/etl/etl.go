// Package etl implements utilities to convert old scrabble score system to a new
package etl

import (
	"strings"
)

// Transform returns new scrabble score system from a given old scrabble score system
func Transform(m map[int][]string) map[string]int {
	var letterScores = make(map[string]int)

	for score, letters := range m {
		for _, letter := range letters {
			letterScores[strings.ToLower(letter)] = score
		}
	}

	return letterScores
}
