package algorithms

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

func FindIfNot[T comparable](collection []T, pred Predicate[T]) int {
	for i, item := range collection {
		if !pred(item) {
			return i
		}
	}
	return -1
}
