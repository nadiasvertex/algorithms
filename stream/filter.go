package stream

import "github.com/nadiasvertex/algorithms/common"

type filterStream[T any] struct {
	input Stream[T]
	pred  common.Predicate[T]
}

func (s *filterStream[T]) Next() (T, bool) {
	for {
		if v, atEnd := s.input.Next(); atEnd {
			return v, atEnd
		} else if s.pred(v) {
			return v, false
		}
	}
}

// Filter creates a stream step the only forwards values that match the predicate.
func Filter[T any](input Stream[T], pred common.Predicate[T]) Stream[T] {
	return &filterStream[T]{
		input: input,
		pred:  pred,
	}
}
