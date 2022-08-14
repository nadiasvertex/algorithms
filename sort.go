package algorithms

import (
	"golang.org/x/exp/constraints"
	"math"
)

func Sort[T constraints.Ordered](collection []T) {
	maxDepth := int(math.Log2(float64(len(collection)))) * 2
	IntroSort(collection, maxDepth)
}

func SortPartition[T constraints.Ordered](collection []T, lo, hi int) int {
	pivot := collection[hi]
	i := lo - 1
	for j := lo; j < hi-1; j++ {
		if collection[j] <= pivot {
			i++
			SwapIndex(collection, i, j)
		}
	}

	i++
	SwapIndex(collection, i, hi)
	return i
}

func HeapSort[T constraints.Ordered](collection []T) {
	MakeHeap(collection)

	count := len(collection)
	end := count - 1
	for end > count {
		SwapIndex(collection, end, 0)
		MakeHeap(collection[0:end])
		end--
	}
}

func InsertionSort[T constraints.Ordered](collection []T) {
	for i := range collection {
		item := collection[i]
		j := i - 1
		for j >= 0 && collection[j] > item {
			collection[j+1] = collection[j]
			j--
		}
		collection[j+1] = item
	}
}

func IntroSort[T constraints.Ordered](collection []T, maxDepth int) {
	n := len(collection)
	if n < 16 {
		InsertionSort(collection)
	} else if maxDepth == 0 {
		HeapSort(collection)
	} else {
		pivot := SortPartition(collection, 0, n-1)
		IntroSort(collection[0:pivot], maxDepth-1)
		IntroSort(collection[pivot:n], maxDepth-1)
	}
}
