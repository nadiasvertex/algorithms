package parallel

import "github.com/nadiasvertex/algorithms/serial"

func Reverse[T any](collection []T) {

	// The parallel version works by chunking half the collection. Each chunk
	// represents a range at the beginning AND the end of the collection. Since
	// reverse is a purely symmetrical operation, the chunks can read and write
	// to their part of the collection without overlapping.

	ProcessMirrorChunks(0, len(collection), func(leftFirst, leftLast, rightFirst, rightLast int) {
		for leftFirst < leftLast {
			serial.SwapIndex(collection, leftFirst, rightLast)
			leftFirst++
			rightLast--
		}
	})

}

func ReverseCopy[T any](collection []T) []T {
	newCollection := make([]T, len(collection))

	// The parallel version works by chunking half the collection. Each chunk
	// represents a range at the beginning AND the end of the collection. Since
	// reverse is a purely symmetrical operation, the chunks can read and write
	// to their part of the collection without overlapping.

	ProcessMirrorChunks(0, len(collection), func(leftFirst, leftLast, rightFirst, rightLast int) {
		for leftFirst < leftLast {
			newCollection[rightLast] = collection[leftFirst]
			newCollection[leftFirst] = collection[rightLast]
			leftFirst++
			rightLast--
		}
	})

	return newCollection
}
