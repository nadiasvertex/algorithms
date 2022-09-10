package stream

type TakeStream[T any] struct {
	input Stream[T]
	taken int
	limit int
}

func (s *TakeStream[T]) Next() (T, bool) {
	if s.taken == s.limit {
		var output T
		return output, true
	}

	s.taken++
	return s.input.Next()
}

func Take[T any](input Stream[T], limit int) Stream[T] {
	return &TakeStream[T]{
		input: input,
		limit: limit,
	}
}
