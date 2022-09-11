package stream

import (
	"reflect"
	"testing"
)

func TestMergeSortInner(t *testing.T) {
	want := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4}
	if got := ToSlice(Take(MergeSort(Iota[int](), Iota[int](), func(value1, value2 int) bool {
		return value1 < value2
	}), 10)); !reflect.DeepEqual(got, want) {
		t.Errorf("ToSlice() = %v, want %v", got, want)
	}
}

func TestMergeSortOuter(t *testing.T) {
	want := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}
	if got := ToSlice(
		MergeSort(
			Take(Iota[int](), 10),
			Take(Iota[int](), 10),
			func(value1, value2 int) bool {
				return value1 < value2
			})); !reflect.DeepEqual(got, want) {
		t.Errorf("ToSlice() = %v, want %v", got, want)
	}
}

func TestMergeSortDifferentSizes(t *testing.T) {
	want := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 7, 8, 9}
	if got := ToSlice(
		MergeSort(
			FromSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}),
			FromSlice([]int{0, 1, 2, 3, 4, 5}),
			func(value1, value2 int) bool {
				return value1 < value2
			})); !reflect.DeepEqual(got, want) {
		t.Errorf("ToSlice() = %v, want %v", got, want)
	}
}
