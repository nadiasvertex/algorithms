package stream

import (
	"github.com/nadiasvertex/algorithms/cnt"
	"github.com/nadiasvertex/algorithms/common"
	"github.com/nadiasvertex/algorithms/serial"
)

type mergeSortStream[T any] struct {
	input1   Stream[T]
	input2   Stream[T]
	holding1 *cnt.Optional[T]
	holding2 *cnt.Optional[T]
	pred     common.BinaryPredicate[T]
}

func (s *mergeSortStream[T]) Next() *cnt.Optional[T] {
	var output *cnt.Optional[T]
	output = cnt.NullOpt[T]()

	// If we are holding nothing from stream1, load its next value.
	if !s.holding1.HasValue() {
		s.holding1 = s.input1.Next()
	}

	// If we are holding nothing from stream2, load its next value.
	if !s.holding2.HasValue() {
		s.holding2 = s.input2.Next()
	}

	// Both streams are empty
	if !s.holding1.HasValue() && !s.holding2.HasValue() {
		return output
	}

	// If both streams have values, pick a value to send.
	if s.holding1.HasValue() && s.holding2.HasValue() {
		if s.pred(s.holding1.Value(), s.holding2.Value()) {
			serial.Swap(&s.holding1, &output)
		} else {
			serial.Swap(&s.holding2, &output)
		}
		return output
	}

	if s.holding1.HasValue() {
		serial.Swap(&s.holding1, &output)
	} else {
		serial.Swap(&s.holding2, &output)
	}

	return output
}

// MergeSort creates a stream that sorts the two streams. The pred function should
// return true if value1 should come before value2. This stream step requires
// that both input streams are already sorted with respect to themselves.
func MergeSort[T any](input1, input2 Stream[T], pred common.BinaryPredicate[T]) Stream[T] {
	return &mergeSortStream[T]{
		input1: input1,
		input2: input2,
		pred:   pred,
	}
}
