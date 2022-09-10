package serial

import "github.com/nadiasvertex/algorithms/common"

// Count counts the number of occurrences of value in collection.
func Count[T comparable](collection []T, value T) int {
	count := 0
	for _, item := range collection {
		if item == value {
			count++
		}
	}
	return count
}

// CountIf counts the number of elements for which pred returns true in collection.
func CountIf[T comparable](collection []T, pred common.Predicate[T]) int {
	count := 0
	for _, item := range collection {
		if pred(item) {
			count++
		}
	}
	return count
}
