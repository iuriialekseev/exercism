// Package raindrops implements tools to convert integers to raindrop sounds
package raindrops

import "strconv"

// Convert converts given integer to raindrop sounds
func Convert(i int) string {
	var result string

	if i%3 == 0 {
		result += "Pling"
	}
	if i%5 == 0 {
		result += "Plang"
	}
	if i%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		result += strconv.Itoa(i)
	}

	return result
}
