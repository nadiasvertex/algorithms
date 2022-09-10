package serial

import (
	"github.com/nadiasvertex/algorithms/common"
)

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
