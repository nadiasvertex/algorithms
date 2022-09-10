package serial

// Fill fills every element of collection with value.
func Fill[T any](collection []T, value T) []T {
	for i := range collection {
		collection[i] = value
	}
	return collection
}
