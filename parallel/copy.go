package parallel

import "github.com/nadiasvertex/algorithms/common"

func Copy[T any](srcCollection []T, dstCollection []T, pred common.Predicate[T]) []T {
	ProcessChunks(0, len(srcCollection), func(first, last int) {
		for i := first; i < last; i++ {
			dstCollection[i] = srcCollection[i]
		}
	})

	return srcCollection
}
