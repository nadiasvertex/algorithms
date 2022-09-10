package stream

import (
	"github.com/nadiasvertex/algorithms/serial"
	"reflect"
	"testing"
)

func TestSkipInner(t *testing.T) {
	want := serial.IotaNWithStart(10, 5)
	got := ToSlice(Take(Skip(Iota[int](), 5), 10))

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Skip() = %v, want %v", got, want)
	}
}

func TestSkipOuter(t *testing.T) {
	want := serial.IotaNWithStart(5, 5)
	got := ToSlice(Skip(Take(Iota[int](), 10), 5))

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Skip() = %v, want %v", got, want)
	}
}

func TestSkipEmpty(t *testing.T) {
	got := ToSlice(Skip(Take(Iota[int](), 5), 5))

	if len(got) > 0 {
		t.Errorf("Skip() = %v, want empty", got)
	}
}

func TestSkipTooMany(t *testing.T) {
	got := ToSlice(Skip(Take(Iota[int](), 5), 10))

	if len(got) > 0 {
		t.Errorf("Skip() = %v, want empty", got)
	}
}
