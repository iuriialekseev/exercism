// Package acronym implements function that abbreviates given string
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate transforms given string to abbreviation
func Abbreviate(s string) string {
	re := regexp.MustCompile("[a-zA-Z']+")
	words := re.FindAllString(s, -1)

	result := ""
	for _, word := range words {
		result += word[:1]
	}

	return strings.ToUpper(result)
}
