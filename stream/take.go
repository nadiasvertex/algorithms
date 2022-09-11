package stream

type takeStream[T any] struct {
	input Stream[T]
	taken int
	limit int
}

func (s *takeStream[T]) Next() (T, bool) {
	if s.taken == s.limit {
		var output T
		return output, true
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
