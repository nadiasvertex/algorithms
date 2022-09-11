package stream

import "github.com/nadiasvertex/algorithms/cnt"

type takeStream[T any] struct {
	input Stream[T]
	taken int
	limit int
}

func (s *takeStream[T]) Next() cnt.Optional[T] {
	if s.taken == s.limit {
		return cnt.NullOpt[T]()
	}

	s.taken++
	return s.input.Next()
}

// Take takes a limited number of items from the stream and then stops.
func Take[T any](input Stream[T], limit int) Stream[T] {
	return &takeStream[T]{
		input: input,
		limit: limit,
	}
}
