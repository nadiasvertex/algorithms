package stream

import (
	"github.com/nadiasvertex/algorithms/serial"
	"reflect"
	"testing"
)

func TestTransformInner(t *testing.T) {
	want := serial.TransformAppend(serial.IotaN[int](10), func(value int) int {
		return value * 2
	})
	if got := ToSlice(Take(Transform(Iota[int](), func(value int) int {
		return value * 2
	}), 10)); !reflect.DeepEqual(got, want) {
		t.Errorf("ToSlice() = %v, want %v", got, want)
	}
}

func TestTransformOuter(t *testing.T) {
	want := serial.TransformAppend(serial.IotaN[int](10), func(value int) int {
		return value * 2
	})
	if got := ToSlice(Transform(Take(Iota[int](), 10), func(value int) int {
		return value * 2
	})); !reflect.DeepEqual(got, want) {
		t.Errorf("ToSlice() = %v, want %v", got, want)
	}
}
