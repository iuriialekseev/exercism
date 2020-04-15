// Package romannumberals implements utilities for converting arabic numbers to romain numerals
package romannumerals

import (
	"errors"
)

type arabicRomanPair struct {
	arabic int
	roman  string
}

var arabicRomanPairs = []arabicRomanPair{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeral converts arabic number to roman numeral
// and returns error if given number is invalid
func ToRomanNumeral(i int) (result string, err error) {
	if i <= 0 || i > 3000 {
		err = errors.New("invalid provided number")
		return
	}

	for _, pair := range arabicRomanPairs {
		for i >= pair.arabic {
			i -= pair.arabic
			result += pair.roman
		}
	}
	return
}
