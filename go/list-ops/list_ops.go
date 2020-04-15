// Package listops implements various list operations
package listops

type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// given two lists, add all items in the second list to the end of the first list
func (list IntList) Append(b IntList) IntList {
	result := list
	result = append(result, b...)

	return result
}

// given a function and a list, return the list of the results of applying given function on all items
func (list IntList) Map(fn unaryFunc) IntList {
	result := IntList{}

	for _, item := range list {
		result = append(result, fn(item))
	}

	return result
}

// given a list, return a list with all the original items, but in reversed order
func (list IntList) Reverse() IntList {
	result := IntList{}

	for i := len(list) - 1; i >= 0; i-- {
		result = append(result, list[i])
	}

	return result
}

// given a series of lists, combine all items in all lists into one flattened list
func (list IntList) Concat(args []IntList) IntList {
	result := list

	for _, arg := range args {
		result = append(result, arg...)
	}

	return result
}

// given a list, return the total number of items within it
func (list IntList) Length() int {
	result := 0

	for range list {
		result++
	}

	return result
}

// given a predicate and a list, return the list of all items for which predicate is True
func (list IntList) Filter(fn predFunc) IntList {
	result := IntList{}

	for _, item := range list {
		if fn(item) {
			result = append(result, item)
		}
	}

	return result
}

// given a function, a list, and initial accumulator,
// fold (reduce) each item into the accumulator from the left using accumulator
func (list IntList) Foldl(fn binFunc, accumulator int) int {
	for _, item := range list {
		accumulator = fn(accumulator, item)
	}

	return accumulator
}

// given a function, a list, and an initial accumulator,
// fold (reduce) each item into the accumulator from the right using accumulator
func (list IntList) Foldr(fn binFunc, accumulator int) int {
	for i := len(list) - 1; i >= 0; i-- {
		accumulator = fn(list[i], accumulator)
	}

	return accumulator
}
