package parallel

import (
	"github.com/nadiasvertex/algorithms/common"
)

func Generate[T any](collection []T, g common.Generator[T]) []T {
	ProcessChunks(0, len(collection), func(first, last int) {
		for i := first; i < last; i++ {
			collection[i] = g(i)
		}
	})

	return collection
}
