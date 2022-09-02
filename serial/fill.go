package serial

func Fill[T any](collection []T, value T) []T {
	for i := range collection {
		collection[i] = value
	}
	return collection
}
