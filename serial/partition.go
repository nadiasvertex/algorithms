package serial

import "github.com/nadiasvertex/algorithms/common"

func Partition[T comparable](collection []T, pred common.Predicate[T]) int {
	first := FindIfNot(collection, pred)
	if first == -1 {
		return len(collection)
	}

	if len(collection) == 1 {
		return first
	}

	for i := first + 1; i < len(collection); i++ {
		if pred(collection[i]) {
			SwapIndex(collection, i, first)
			first++
		}
	}

	return first
}
