// Package etl implements utilities to convert old scrabble score system to a new
package etl

import (
	"strings"
)

// Transform returns new scrabble score system from a given old scrabble score system
func Transform(input map[int][]string) (output map[string]int) {
	output = make(map[string]int)

	for score, letters := range input {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = score
		}
	}
	return
}
