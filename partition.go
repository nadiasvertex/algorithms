package algorithms

func Partition[T comparable](collection []T, pred Predicate[T]) int {
	first := FindIfNot(collection, pred)
	if first == -1 {
		return len(collection)
	}

	for i := first + 1; first < len(collection); i++ {
		if pred(collection[i]) {
			SwapIndex(collection, i, first)
			first++
		}
	}

	return first
}
