// Package triangle implements function that determines triangle kind based on given sides
package triangle

import (
	"math"
)

type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

var kinds = []Kind{NaT, Equ, Iso, Sca}

// KindFromSides accepts sides of a triangle and returns one of four triangle kinds
func KindFromSides(a, b, c float64) Kind {
	sides := []float64{a, b, c}
	uniqueSides := make(map[float64]bool)
	perimeter := a + b + c

	for _, side := range sides {
		if math.IsNaN(side) || side <= 0 || math.IsInf(side, 1) || perimeter-side < side {
			return NaT
		}
	}

	for _, side := range sides {
		uniqueSides[side] = true
	}

	return kinds[len(uniqueSides)]
}
