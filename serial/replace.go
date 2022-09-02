package serial

import "github.com/nadiasvertex/algorithms/common"

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
