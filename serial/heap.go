package serial

import (
	"golang.org/x/exp/constraints"
)

func MakeHeap[T constraints.Ordered](collection []T) {
	size := len(collection)
	start := size/2 - 1 // last non leaf
	for start >= 0 {
		siftDown(collection, start, size-1)
		start = start - 1
	}
}

func siftDown[T constraints.Ordered](collection []T, start, end int) {
	root := start
	for root*2+1 <= end {
		lchild := root*2 + 1 // subscript of left child
		swap := root
		if collection[swap] < collection[lchild] {
			swap = lchild // should swap root & left child
		}
		if lchild+1 <= end && collection[swap] < collection[lchild+1] {
			swap = lchild + 1 // should swap root & right child
		}
		if swap != root { // if one of children is greater
			SwapIndex(collection, root, swap) // swap root & larger child
			root = swap
		} else {
			return
		}
	}
}
