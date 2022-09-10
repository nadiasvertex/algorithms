package stream

import (
	"reflect"
	"testing"
)

func TestToSlice(t *testing.T) {
	want := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	if got := ToSlice(Take(Iota[int](), 10)); !reflect.DeepEqual(got, want) {
		t.Errorf("ToSlice() = %v, want %v", got, want)
	}
}
