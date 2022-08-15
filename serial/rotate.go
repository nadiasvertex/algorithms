package serial

func Rotate[T any](collection []T, nFirst int) int {
	return RotateRange(collection, 0, nFirst, len(collection))
}

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
