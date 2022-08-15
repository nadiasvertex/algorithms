package serial

// Unique eliminates all except the first element from every consecutive group
// of equivalent elements from the collection.
func Unique[T comparable](collection []T) []T {
	first := 0
	last := len(collection)

	if first == last {
		return nil
	}

	result := first + 1
	for ; first != last; first++ {
		if !(collection[result] == collection[first]) && collection[result+1] != collection[first] {
			result++
			collection[result] = collection[first]
		}
	}

	return collection[0 : result+1]
}
