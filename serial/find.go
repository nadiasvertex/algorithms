package serial

import "github.com/nadiasvertex/algorithms/common"

// Find locates the index of first occurrence of value in collection. If the
// value is not present it returns -1.
func Find[T comparable](collection []T, value T) int {
	for i, item := range collection {
		if item == value {
			return i
		}
	}
	return -1
}

// FindNot locates the index of first element in collection that is not equal to
// value. If the value is not present it returns -1.
func FindNot[T comparable](collection []T, value T) int {
	for i, item := range collection {
		if item != value {
			return i
		}
	}
	return -1
}

// FindIf locates the index of the first element for which pred returns true.
// If no element returns true for pred the function returns -1.
func FindIf[T comparable](collection []T, pred common.Predicate[T]) int {
	for i, item := range collection {
		if pred(item) {
			return i
		}
	}
	return -1
}

// FindIfNot locates the index of the first element for which pred returns false.
// If no element returns false for pred the function returns -1.
func FindIfNot[T comparable](collection []T, pred common.Predicate[T]) int {
	for i, item := range collection {
		if !pred(item) {
			return i
		}
	}
	return -1
}
