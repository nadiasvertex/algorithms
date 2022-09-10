package serial

import "github.com/nadiasvertex/algorithms/common"

// Contains returns true if value is in collection.
func Contains[T comparable](collection []T, value T) bool {
	return Find(collection, value) >= 0
}

// ContainsIf returns true if any element in collection returns true for pred.
func ContainsIf[T comparable](collection []T, pred common.Predicate[T]) bool {
	return FindIf(collection, pred) >= 0
}
