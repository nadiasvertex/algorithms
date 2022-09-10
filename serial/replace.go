package serial

import "github.com/nadiasvertex/algorithms/common"

// Replace replaces all elements that are equal to old_value with new_value.
// Returns the number of replacements made.
func Replace[T comparable](collection []T, old_value, new_value T) int {
	count := 0
	for i, item := range collection {
		if item == old_value {
			collection[i] = new_value
			count++
		}
	}
	return count
}

// ReplaceIf replaces all elements where pred is true with new_value.
// Returns the number of replacements made.
func ReplaceIf[T comparable](collection []T, pred common.Predicate[T], new_value T) int {
	count := 0
	for i, item := range collection {
		if pred(item) {
			collection[i] = new_value
			count++
		}
	}
	return count
}
