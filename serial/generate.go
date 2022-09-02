package serial

import "github.com/nadiasvertex/algorithms/common"

func Generate[T any](collection []T, g common.Generator[T]) []T {
	for i := range collection {
		collection[i] = g(i)
	}
	return collection
}
