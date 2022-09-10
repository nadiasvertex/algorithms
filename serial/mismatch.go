package serial

// Mismatch returns the first mismatching pair of elements from two slices. The
// return value is the first index in collection1 that does not match the same
// index in collection2.
func Mismatch[T comparable](collection1, collection2 []T) int {
	for i := 0; i < len(collection1); i++ {
		if collection1[i] != collection2[i] {
			return i
		}
	}

	return len(collection1)
}
