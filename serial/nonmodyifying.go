package serial

import (
	"github.com/nadiasvertex/algorithms/common"
	"golang.org/x/exp/constraints"
)

// Max returns the greater of the two values, or the first value if they are
// the same.
func Max[T constraints.Ordered](v1, v2 T) T {
	if v1 >= v2 {
		return v1
	}
	return v2
}

// Min returns the lesser of the two values, or the first value if they are the
// same.
func Min[T constraints.Ordered](v1, v2 T) T {
	if v1 <= v2 {
		return v1
	}
	return v2
}

// AllOf returns true if all the elements in the collection are equal to value.
func AllOf[T comparable](collection []T, value T) bool {
	for _, item := range collection {
		if item != value {
			return false
		}
	}
	return true
}

// AllOfIf returns true if all the elements in the collection return true for pred.
func AllOfIf[T comparable](collection []T, pred common.Predicate[T]) bool {
	for _, item := range collection {
		if !pred(item) {
			return false
		}
	}
	return true
}

// NoneOf returns true if none of the elements in the collection are equal to value.
func NoneOf[T comparable](collection []T, value T) bool {
	for _, item := range collection {
		if item == value {
			return false
		}
	}
	return true
}

// NoneOfIf returns true if none of the elements in the collection return true for pred.
func NoneOfIf[T comparable](collection []T, pred common.Predicate[T]) bool {
	for _, item := range collection {
		if pred(item) {
			return false
		}
	}
	return true
}

// AnyOf returns true if any of the elements in the collection are equal to value.
func AnyOf[T comparable](collection []T, value T) bool {
	for _, item := range collection {
		if item == value {
			return true
		}
	}
	return false
}

// AnyOfIf returns true if any of the elements in the collection return true for pred.
func AnyOfIf[T comparable](collection []T, pred common.Predicate[T]) bool {
	for _, item := range collection {
		if pred(item) {
			return true
		}
	}
	return false
}

// Mismatch returns the first mismatching pair of elements from two slices. The
// return value is the first index in collection1 that does not match the same
// index in collection2.
func Mismatch[T comparable](collection1, collection2 []T) int {
	for i := 0; i < len(collection1); i++ {
		if collection1[i] != collection2[i] {
			return i
		}
	}

	return len(collection1)
}
