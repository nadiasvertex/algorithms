package parallel

import (
	"github.com/nadiasvertex/algorithms/common"
	"sync/atomic"
)

func Replace[T comparable](collection []T, oldValue, newValue T) int {
	var count atomic.Int64

	ProcessChunks(0, len(collection), func(first, last int) {
		for i := first; i < last; i++ {
			item := collection[i]
			if item == oldValue {
				collection[i] = newValue
				count.Add(1)
			}
		}
	})

	return int(count.Load())
}

func ReplaceIf[T comparable](collection []T, pred common.Predicate[T], newValue T) int {
	var count atomic.Int64

	ProcessChunks(0, len(collection), func(first, last int) {
		for i := first; i < last; i++ {
			item := collection[i]
			if pred(item) {
				collection[i] = newValue
				count.Add(1)
			}
		}
	})

	return int(count.Load())
}
