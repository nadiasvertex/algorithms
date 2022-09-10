package stream

type SkipStream[T any] struct {
	input   Stream[T]
	skipped int
	limit   int
}

func (s *SkipStream[T]) Next() (T, bool) {
	for ; s.skipped < s.limit; s.skipped++ {
		if v, atEnd := s.input.Next(); atEnd {
			return v, atEnd
		}
	}
	return s.input.Next()
}

func Skip[T any](input Stream[T], limit int) Stream[T] {
	return &SkipStream[T]{
		input: input,
		limit: limit,
	}
}
