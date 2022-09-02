package serial

import "github.com/nadiasvertex/algorithms/common"

// Copy copies all items in srcCollection to dstCollection.
func Copy[T comparable](srcCollection []T, dstCollection []T) {
	for i, item := range srcCollection {
		dstCollection[i] = item
	}
}

// CopyIf copies items from srcCollection that match the given predicate. dstCollection
// must have space for at most the same number of items as srcCollection. If you
// know a priori how big the destination collection must be, this will more efficient than
// CopyAppendIf.
func CopyIf[T comparable](srcCollection []T, dstCollection []T, pred common.Predicate[T]) int {
	output := 0
	for _, item := range srcCollection {
		if pred(item) {
			dstCollection[output] = item
			output++
		}
	}
	return output
}

// CopyAppendIf copies items from srcCollection that match the given predicate. The
// items are copied to a new collection. The return value is a collection of items
// that match the predicate
func CopyAppendIf[T comparable](srcCollection []T, pred common.Predicate[T]) []T {
	var dstCollection []T
	for _, item := range srcCollection {
		if pred(item) {
			dstCollection = append(dstCollection, item)
		}
	}
	return dstCollection
}
