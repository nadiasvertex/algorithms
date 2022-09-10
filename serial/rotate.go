package serial

// Rotate swaps the elements in the collection in such a way that the
// element nFirst becomes the first element of the new collection and nFirst - 1
// becomes the last element.
func Rotate[T any](collection []T, nFirst int) int {
	return RotateRange(collection, 0, nFirst, len(collection))
}

// RotateRange swaps the elements in the range [first, last) in such a way that
// the element nFirst becomes the first element of the new range and nFirst - 1
// becomes the last element. A precondition of this function is that
// [first, nFirst) and [nFirst, last) are valid ranges.
func RotateRange[T any](collection []T, first, nFirst, last int) int {
	if nFirst == first {
		return len(collection)
	}

	if nFirst == last {
		return 0
	}

	read := nFirst
	write := first
	nextRead := first

	for read != last {
		if write == nextRead {
			nextRead = read
		}
		SwapIndex(collection, write, read)
		write++
		read++
	}

	RotateRange(collection, write, nextRead, last)
	return write
}
