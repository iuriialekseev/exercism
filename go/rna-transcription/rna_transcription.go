// Package strand implements DNA to RNA transpiler
package strand

import (
	"strings"
)

var DNAtoRNAMappings = map[string]string{
	`G`: `C`,
	`C`: `G`,
	`T`: `A`,
	`A`: `U`,
}

// ToRNA accepts DNA string and returns it's RNA analogue
func ToRNA(dna string) string {
	rna := strings.Split(dna, "")

	for i, nucleotide := range rna {
		rna[i] = DNAtoRNAMappings[nucleotide]
	}

	return strings.Join(rna, "")
}
