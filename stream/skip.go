package stream

type skipStream[T any] struct {
	input   Stream[T]
	skipped int
	limit   int
}

func (s *skipStream[T]) Next() (T, bool) {
	for ; s.skipped < s.limit; s.skipped++ {
		if v, atEnd := s.input.Next(); atEnd {
			return v, atEnd
		}
	}
	return s.input.Next()
}

// Skip skips a certain number of items from the input stream and delivers the
// remainder.
func Skip[T any](input Stream[T], limit int) Stream[T] {
	return &skipStream[T]{
		input: input,
		limit: limit,
	}
}
