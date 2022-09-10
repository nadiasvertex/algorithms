package serial

import "golang.org/x/exp/constraints"

// Max returns the greater of the two values, or the first value if they are
// the same.
func Max[T constraints.Ordered](v1, v2 T) T {
	if v1 >= v2 {
		return v1
	}
	return v2
}

// Min returns the lesser of the two values, or the first value if they are the
// same.
func Min[T constraints.Ordered](v1, v2 T) T {
	if v1 <= v2 {
		return v1
	}
	return v2
}

// MaxElement finds the index of the first occurrence of the largest item in the
// collection.
func MaxElement[T constraints.Ordered](collection []T) int {
	max := 0
	for i := range collection {
		if collection[i] > collection[max] {
			max = i
		}
	}
	return max
}

// MinElement finds the index of the first occurrence of the smallest item in the
// collection.
func MinElement[T constraints.Ordered](collection []T) int {
	min := 0
	for i := range collection {
		if collection[i] < collection[min] {
			min = i
		}
	}
	return min
}

// MinMaxElement finds the index of the first occurrence of the smallest and largest items in the
// collection. Returns a tuple of (min_index, max_index).
func MinMaxElement[T constraints.Ordered](collection []T) (int, int) {
	min := 0
	max := 0
	for i := range collection {
		if collection[i] < collection[min] {
			min = i
		}
		if collection[i] > collection[max] {
			max = i
		}
	}
	return min, max
}
