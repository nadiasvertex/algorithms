package serial

import (
	"golang.org/x/exp/constraints"
)

// MakeHeap turns the collection into a heap, in place.
func MakeHeap[T constraints.Ordered](collection []T) []T {
	size := len(collection)
	start := size/2 - 1 // last non leaf
	for start >= 0 {
		siftDown(collection, start, size-1)
		start = start - 1
	}
	return collection
}

// siftDown performs the sift down heap establishing operation to the range
// [start, end) in collection.
func siftDown[T constraints.Ordered](collection []T, start, end int) {
	root := start
	for root*2+1 <= end {
		leftChild := root*2 + 1 // subscript of left child
		swap := root
		if collection[swap] < collection[leftChild] {
			swap = leftChild // should swap root & left child
		}
		if leftChild+1 <= end && collection[swap] < collection[leftChild+1] {
			swap = leftChild + 1 // should swap root & right child
		}
		if swap != root { // if one of children is greater
			SwapIndex(collection, root, swap) // swap root & larger child
			root = swap
		} else {
			return
		}
	}
}
