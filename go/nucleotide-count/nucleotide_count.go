// Package dna implements tools to calculate
// number of different nucleotides in dna strand
package dna

import "errors"

type Histogram map[rune]int
type DNA []rune

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for _, nucleotide := range d {
		_, ok := h[nucleotide]
		if !ok {
			return h, errors.New("contains invalid nucleotide")
		}

		h[nucleotide]++
	}

	return h, nil
}
