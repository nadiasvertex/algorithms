package serial

import (
	"golang.org/x/exp/constraints"
	"math"
)

// Sort sorts the collection in ascending order. This version of sort uses the
// best non-stable algorithm known to the library. Currently, that is introsort,
// but may change across releases.
func Sort[T constraints.Ordered](collection []T) {
	maxDepth := int(math.Log2(float64(len(collection)))) * 2
	IntroSortWithDepth(collection, maxDepth)
}

// sortPartition performs a partition operation on collection suitable for
// quicksort-like algorithms. This may differ from Partition, so it is
// implemented here and is not exposed.
func sortPartition[T constraints.Ordered](collection []T, lo, hi int) int {
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

// HeapSort sorts the collection using the heap sort algorithm.
func HeapSort[T constraints.Ordered](collection []T) {
	MakeHeap(collection)
	size := len(collection)
	last := size - 1 // subscript of last item
	for last > 0 {
		SwapIndex(collection, last, 0)
		last = last - 1
		siftDown(collection, 0, last)
	}
}

// InsertionSort sorts the collection using the insertion sort algorithm.
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

// IntroSort sorts the collection in ascending order using the introsort
// algorithm and a default depth for recursion bail-out.
func IntroSort[T constraints.Ordered](collection []T) {
	maxDepth := int(math.Log2(float64(len(collection)))) * 2
	IntroSortWithDepth(collection, maxDepth)
}

// IntroSortWithDepth sorts the collection using the introsort algorithm. The maxDepth
// parameter is an indication of how many deep the recursion can go before
// bailing out and falling back to heap sort. This prevents the worst case
// behavior of quicksort. Normally that should be set to int(math.Log2(float64(len(collection)))) * 2,
// which is what the IntroSort function sets it to. It shouldn't be larger than that,
// but it may be less if you want to fall back to heap sort sooner.
func IntroSortWithDepth[T constraints.Ordered](collection []T, maxDepth int) {
	n := len(collection)
	if n < 16 {
		InsertionSort(collection)
	} else if maxDepth == 0 {
		HeapSort(collection)
	} else {
		pivot := sortPartition(collection, 0, n-1)
		IntroSortWithDepth(collection[0:pivot], maxDepth-1)
		IntroSortWithDepth(collection[pivot:n], maxDepth-1)
	}
}
