package stream

import "github.com/nadiasvertex/algorithms/parallel"

// ChunkOut drains a Stream and provides its output as an array of chunks.
func ChunkOut[T any](input Stream[T], size int) [][]T {
	var output [][]T
	chunk := make([]T, size)
	currentSize := 0
	for {
		if v := input.Next(); !v.HasValue() {
			if currentSize > 0 {
				output = append(output, chunk[0:currentSize])
			}
			break
		} else {
			chunk[currentSize] = v.Value()
			currentSize++
			if currentSize == size {
				output = append(output, chunk)
				chunk = make([]T, size)
				currentSize = 0
			}
		}
	}

	return output
}

// ChunkIn provides an array of streams, each one draining a chunk of the input
// slice.
func ChunkIn[T any](input []T, size int) []Stream[T] {
	var chunks []Stream[T]
	for _, chunk := range parallel.Chunks(0, len(input), size) {
		chunks = append(chunks, FromSlice(input[chunk.First:chunk.Last]))
	}
	return chunks
}
