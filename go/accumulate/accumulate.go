// Package accumulate implements map-like utilities
package accumulate

type converter func(string) string

// Accumulate accepts a collection and an operation to perform on it's members
// and returns a new collection containing the result of applying that operation
func Accumulate(s []string, c converter) []string {
	var result []string

	for _, item := range s {
		result = append(result, c(item))
	}

	return result
}
