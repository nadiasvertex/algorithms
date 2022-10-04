package stream

import "github.com/nadiasvertex/algorithms/cnt"

type skipStream[T any] struct {
	input   Stream[T]
	skipped int
	limit   int
}

func (s *skipStream[T]) Next() *cnt.Optional[T] {
	for ; s.skipped < s.limit; s.skipped++ {
		if v := s.input.Next(); !v.HasValue() {
			return v
		}
	}
	return s.input.Next()
}

// Skip skips a certain number of items from the input Stream and delivers the
// remainder.
func Skip[T any](input Stream[T], limit int) Stream[T] {
	return &skipStream[T]{
		input: input,
		limit: limit,
	}
}
