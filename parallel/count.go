package parallel

import (
	"github.com/nadiasvertex/algorithms/serial"
	"sync/atomic"
)

func Count[T comparable](collection []T, value T) int {
	var count atomic.Int64

	ProcessChunks(0, len(collection), func(first, last int) {
		localCount := serial.Count(collection[first:last], value)
		count.Add(int64(localCount))
	})

	return int(count.Load())
}
