package algorithms

func Contains[T comparable](collection []T, value T) bool {
	return Find(collection, value) >= 0
}

func ContainsIf[T comparable](collection []T, pred Predicate[T]) bool {
	return FindIf(collection, pred) >= 0
}
