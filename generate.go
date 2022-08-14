package algorithms

func Generate[T comparable](collection []T, g Generator[T]) int {
	count := 0
	for i := range collection {
		collection[i] = g()
	}
	return count
}
