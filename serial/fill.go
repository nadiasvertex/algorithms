package serial

func Fill[T comparable](collection []T, value T) int {
	count := 0
	for i := range collection {
		collection[i] = value
	}
	return count
}
