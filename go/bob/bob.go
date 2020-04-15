// Package bob implements function that represents bob's answers
package bob

import (
	"strings"
	"unicode"
)

// Hey accepts phrase and returns bob's answer to that phrase
func Hey(remark string) string {
	trimmedRemark := strings.TrimSpace(remark)

	isSilence := trimmedRemark == ""
	hasLetters := strings.IndexFunc(trimmedRemark, unicode.IsLetter) >= 0
	isYelling := hasLetters && strings.ToUpper(trimmedRemark) == trimmedRemark
	isQuestion := strings.HasSuffix(trimmedRemark, "?")

	switch {
	case isSilence:
		return "Fine. Be that way!"
	case isYelling && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isYelling:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	default:
		return "Whatever."
	}
}
