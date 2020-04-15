// Package scrabble implements function to calculate scrabble score of a string
package scrabble

import "strings"

type LettersScore struct {
	letters string
	score   int
}

var ScrabbleScoreCard = []LettersScore{
	{"AEIOULNRST", 1},
	{"DG", 2},
	{"BCMP", 3},
	{"FHVWY", 4},
	{"K", 5},
	{"JX", 8},
	{"QZ", 10},
}

// Score returns scrabble score of a given string
func Score(s string) (score int) {
	upperStr := strings.ToUpper(s)

	for _, letter := range upperStr {
		for _, letterScore := range ScrabbleScoreCard {
			if strings.ContainsRune(letterScore.letters, letter) {
				score += letterScore.score
			}
		}
	}
	return
}
