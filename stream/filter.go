package stream

import "github.com/nadiasvertex/algorithms/common"

type FilterStream[T any] struct {
	input Stream[T]
	pred  common.Predicate[T]
}

func (s *FilterStream[T]) Next() (T, bool) {
	for {
		if v, atEnd := s.input.Next(); atEnd {
			return v, atEnd
		} else if s.pred(v) {
			return v, false
		}
	}
}

func Filter[T any](input Stream[T], pred common.Predicate[T]) Stream[T] {
	return &FilterStream[T]{
		input: input,
		pred:  pred,
	}
}
