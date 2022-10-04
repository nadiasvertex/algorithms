package stream

import (
	"github.com/nadiasvertex/algorithms/cnt"
	"github.com/nadiasvertex/algorithms/common"
)

type filterStream[T any] struct {
	input Stream[T]
	pred  common.Predicate[T]
}

func (s *filterStream[T]) Next() *cnt.Optional[T] {
	for {
		if v := s.input.Next(); !v.HasValue() || s.pred(v.Value()) {
			return v
		}
	}
}

// Filter creates a Stream step the only forwards values that match the predicate.
func Filter[T any](input Stream[T], pred common.Predicate[T]) Stream[T] {
	return &filterStream[T]{
		input: input,
		pred:  pred,
	}
}
