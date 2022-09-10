package serial

// Reverse reverses the order of the elements in the collection. Behaves as if
// applying SwapIndex to every pair of iterators first+i, (last-i) - 1 for each
// non-negative i < (last-first)/2. Returns the collection passed in.
func Reverse[T any](collection []T) []T {
	first := 0
	last := len(collection) - 1
	for first < last {
		SwapIndex(collection, first, last)
		first++
		last--
	}
	return collection
}

// ReverseCopy performs the same operation as Reverse, but does it on a copy
// instead of doing it in place. Returns a reversed copy of the input.
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
