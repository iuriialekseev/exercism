// Package hamming implements function that calculates
// hamming difference between two strings
package hamming

import "errors"

// Distance accepts two strings and returns hamming difference between them
func Distance(a, b string) (int, error) {
	distance := 0
	aSlice, bSlice := []byte(a), []byte(b)

	if len(aSlice) != len(bSlice) {
		return 0, errors.New("length is different")
	}

	for i, char := range aSlice {
		if char != bSlice[i] {
			distance++
		}
	}

	return distance, nil
}
