package parallel

import (
	"github.com/nadiasvertex/algorithms/serial"
	"sync/atomic"
)

func Find[T comparable](collection []T, value T) int {
	var index atomic.Int64
	index.Store(int64(len(collection)))

	ProcessChunks(0, len(collection), func(first, last int) {
		i := serial.Find(collection[first:last], value)
		if i != -1 {
			for {
				globalFound := index.Load()
				localFound := int64(i)
				if localFound < globalFound {
					if !index.CompareAndSwap(globalFound, localFound) {
						continue
					}
				}
				break
			}
		}
	})

	return int(index.Load())
}
