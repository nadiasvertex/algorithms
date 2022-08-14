package algorithms

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](v1, v2 T) T {
	if v1 > v2 {
		return v1
	}
	return v2
}

func Find[T comparable](collection []T, value T) int {
	for i, item := range collection {
		if item == value {
			return i
		}
	}
	return -1
}

func FindIf[T comparable](collection []T, pred Predicate[T]) int {
	for i, item := range collection {
		if pred(item) {
			return i
		}
	}
	return -1
}

func Contains[T comparable](collection []T, value T) bool {
	return Find(collection, value) >= 0
}

func ContainsIf[T comparable](collection []T, pred Predicate[T]) bool {
	return FindIf(collection, pred) >= 0
}

func AllOf[T comparable](collection []T, value T) bool {
	for _, item := range collection {
		if item != value {
			return false
		}
	}
	return true
}

func AllOfIf[T comparable](collection []T, pred Predicate[T]) bool {
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

func NoneOfIf[T comparable](collection []T, pred Predicate[T]) bool {
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

func AnyOfIf[T comparable](collection []T, pred Predicate[T]) bool {
	for _, item := range collection {
		if pred(item) {
			return true
		}
	}
	return false
}

func Count[T comparable](collection []T, value T) int {
	count := 0
	for _, item := range collection {
		if item == value {
			count++
		}
	}
	return count
}

func CountIf[T comparable](collection []T, pred Predicate[T]) int {
	count := 0
	for _, item := range collection {
		if pred(item) {
			count++
		}
	}
	return count
}

func Mismatch[T comparable](collection1, collection2 []T) int {
	for i := 0; i < len(collection1); i++ {
		if collection1[i] != collection2[i] {
			return i
		}
	}

	return len(collection1)
}
