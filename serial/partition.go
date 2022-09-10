package serial

import "github.com/nadiasvertex/algorithms/common"

// Partition reorders the elements in the range [first, last) in such a way
// that all elements for which the predicate p returns true precede the elements
// for which predicate p returns false. Relative order of the elements is not
// preserved. The return value is the index to the first element of the second
// group.
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
