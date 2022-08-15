package serial

import "golang.org/x/exp/rand"

// RandomShuffle Reorders the elements in the given collection such that each
// possible permutation of those elements has equal probability of appearance.
func RandomShuffle[T comparable](collection []T) {
	n := len(collection) - 1
	for i := range collection {
		SwapIndex(collection, i, i+rand.Intn(n))
	}
}
