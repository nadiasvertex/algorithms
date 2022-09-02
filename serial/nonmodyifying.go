package serial

import (
	"github.com/nadiasvertex/algorithms/common"
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](v1, v2 T) T {
	if v1 > v2 {
		return v1
	}
	return v2
}

func AllOf[T comparable](collection []T, value T) bool {
	for _, item := range collection {
		if item != value {
			return false
		}
	}
	return true
}

func AllOfIf[T comparable](collection []T, pred common.Predicate[T]) bool {
	for _, item := range collection {
		if !pred(item) {
			return false
		}
	}
	return true
}

func NoneOf[T comparable](collection []T, value T) bool {
	for _, item := range collection {
		if item == value {
			return false
		}
	}
	return true
}

func NoneOfIf[T comparable](collection []T, pred common.Predicate[T]) bool {
	for _, item := range collection {
		if pred(item) {
			return false
		}
	}
	return true
}

func AnyOf[T comparable](collection []T, value T) bool {
	for _, item := range collection {
		if item == value {
			return true
		}
	}
	return false
}

func AnyOfIf[T comparable](collection []T, pred common.Predicate[T]) bool {
	for _, item := range collection {
		if pred(item) {
			return true
		}
	}
	return false
}

func Mismatch[T comparable](collection1, collection2 []T) int {
	for i := 0; i < len(collection1); i++ {
		if collection1[i] != collection2[i] {
			return i
		}
	}

	return len(collection1)
}
