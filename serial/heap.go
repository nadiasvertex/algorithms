package serial

import (
	"golang.org/x/exp/constraints"
	"math"
)

func MakeHeap[T constraints.Ordered](collection []T) {
	count := len(collection)
	end := 1
	for end < count {
		siftUp(collection, 0, end)
		end++
	}
}

func parentIndex(i int) int {
	return int(math.Floor(float64(i-1) / 2.0))
}

func siftUp[T constraints.Ordered](collection []T, start, end int) {
	child := end
	for child > start {
		parent := parentIndex(child)
		if collection[parent] < collection[child] {
			SwapIndex(collection, parent, child)
			child = parent
		} else {
			return
		}
	}
}
