// Package protein implements utilities to translate RNA sequences into proteins
package protein

import (
	"errors"
)

var (
	// ErrInvalidBase for not existing codons
	ErrInvalidBase = errors.New("codon not exists")
	// ErrStop for terminating codons in rna sequence
	ErrStop = errors.New("terminating codon reached")
)

const codonLength = 3

// FromRNA returns proteins from given rna
// and any encountered in transcription process errors
func FromRNA(rna string) (proteins []string, err error) {
	for len(rna) > 0 {
		codon := rna[:codonLength]
		rna = rna[codonLength:]

		protein, err := FromCodon(codon)
		if err == ErrStop {
			return proteins, nil
		}
		if err == ErrInvalidBase {
			return proteins, err
		}

		proteins = append(proteins, protein)
	}
	return
}

// FromCodon returns matching protein from given codon
func FromCodon(codon string) (protein string, err error) {
	switch codon {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	case "UAA", "UAG", "UGA":
		err = ErrStop
	default:
		err = ErrInvalidBase
	}
	return
}
