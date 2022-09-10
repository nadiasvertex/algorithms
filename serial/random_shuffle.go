package serial

import "golang.org/x/exp/rand"

// RandomShuffle reorders the elements in the given collection such that each
// possible permutation of those elements has equal probability of appearance.
func RandomShuffle[T comparable](collection []T) []T {
	n := len(collection)
	for i := range collection {
		SwapIndex(collection, i, i+rand.Intn(n-i))
	}
	return collection
}
