package stream

import (
	"github.com/nadiasvertex/algorithms/serial"
	"testing"
)

func TestForEach(t *testing.T) {
	found := make([]bool, 10)

	ForEach(Take(Iota[int](), 10), func(value int) {
		found[value] = true
	})

	if !serial.AllOf(found, true) {
		t.Errorf("ForEach() expected to take 10 items, got %v", found)
	}
}
