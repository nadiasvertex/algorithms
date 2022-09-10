package serial

import "github.com/nadiasvertex/algorithms/common"

// Remove removes all elements that are equal to value. Returns a new slice
// with items removed.
func Remove[T comparable](collection []T, value T) []T {
	pivot := Partition(collection, func(item T) bool {
		return item != value
	})

	if pivot == len(collection) {
		return nil
	}

	return collection[0:pivot]
}

// RemoveIf removes all elements where pred returns true. Returns a new slice
// with items removed.
func RemoveIf[T comparable](collection []T, pred common.Predicate[T]) []T {
	pivot := Partition(collection, func(item T) bool {
		return !pred(item)
	})

	if pivot == len(collection) {
		return nil
	}

	return collection[0:pivot]
}
