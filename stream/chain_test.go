package stream

import (
	"github.com/nadiasvertex/algorithms/serial"
	"testing"
)

func Test_chainFilter(t *testing.T) {
	even := func(v int) bool {
		return v%2 == 0
	}

	results := Monotonic[int]().Filter(even).Take(10).ToSlice()

	if len(results) < 10 {
		t.Errorf("Expected 10 numbers, got %d", len(results))
	}

	if !serial.AllOfIf(results, even) {
		t.Errorf("Expected 10 even numbers, got %v", results)
	}
}

func Test_chainOneOrNone(t *testing.T) {
	even := func(v int) bool {
		return v%2 == 0
	}

	result := Monotonic[int]().Filter(even).OneOrNone()

	if !result.HasValue() {
		t.Error("Expected a value, got none.")
	}

	if result.Value()%2 != 0 {
		t.Errorf("Expected an even number, got %v", result.Value())
	}
}
