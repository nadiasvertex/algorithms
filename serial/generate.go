package serial

import "github.com/nadiasvertex/algorithms/common"

// Generate fills the collection with the results of the generator function.
func Generate[T any](collection []T, g common.Generator[T]) []T {
	for i := range collection {
		collection[i] = g(i)
	}
	return collection
}
