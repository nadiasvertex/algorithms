package parallel

func Fill[T any](collection []T, value T) []T {
	ProcessChunks(0, len(collection), func(first, last int) {
		for i := first; i < last; i++ {
			collection[i] = value
		}
	})

	return collection
}
