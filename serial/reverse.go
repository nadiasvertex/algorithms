package serial

func Reverse[T any](collection []T) {
	first := 0
	last := len(collection) - 1
	for first < last {
		SwapIndex(collection, first, last)
		first++
		last--
	}
}

func ReverseCopy[T any](collection []T) []T {
	newCollection := make([]T, len(collection))
	last := len(collection) - 1
	first := 0
	for first < len(collection) {
		newCollection[first] = collection[last]
		first++
		last--
	}
	return newCollection
}
