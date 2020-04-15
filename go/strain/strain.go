// Package strain implements select- and reject-like utilities
package strain

type intPredicate func(int) bool
type listsPredicate func([]int) bool
type stringPredicate func(string) bool

type Ints []int
type Lists [][]int
type Strings []string

// Keep returns slice if Ints that pass predicate
func (i Ints) Keep(p intPredicate) (o Ints) {
	for _, item := range i {
		if p(item) {
			o = append(o, item)
		}
	}
	return
}

// Keep returns slice if Ints that don't pass predicate
func (i Ints) Discard(p intPredicate) (o Ints) {
	for _, item := range i {
		if !p(item) {
			o = append(o, item)
		}
	}
	return
}

// Keep returns slice if Strings that pass predicate
func (i Strings) Keep(p stringPredicate) (o Strings) {
	for _, item := range i {
		if p(item) {
			o = append(o, item)
		}
	}
	return
}

// Keep returns slice if Lists that pass predicate
func (i Lists) Keep(p listsPredicate) (o Lists) {
	for _, item := range i {
		if p(item) {
			o = append(o, item)
		}
	}
	return
}
