// Package protein implements utilities to translate RNA sequences into proteins
package protein

import (
	"errors"
)

var codonProteins = map[string]string{
	"AUG": "Methionine",
	"UAA": "STOP",
	"UAC": "Tyrosine",
	"UAG": "STOP",
	"UAU": "Tyrosine",
	"UCA": "Serine",
	"UCC": "Serine",
	"UCG": "Serine",
	"UCU": "Serine",
	"UGA": "STOP",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UGU": "Cysteine",
	"UUA": "Leucine",
	"UUC": "Phenylalanine",
	"UUG": "Leucine",
	"UUU": "Phenylalanine",
}

// ErrInvalidBase for not existing codons
var ErrInvalidBase = errors.New("codon not exists")

// ErrStop for terminating codons in rna sequence
var ErrStop = errors.New("terminating codon reached")

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
	protein, ok := codonProteins[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}
	return
}
