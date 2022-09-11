package stream

import "golang.org/x/exp/constraints"

type generateStream[T any] struct {
	g func() (T, bool)
}

func (s *generateStream[T]) Next() (T, bool) {
	return s.g()
}

// Generate creates a stream that generates user-defined values. When the
// generator wants to terminate the stream, it should return true as it's
// second resut argument.
func Generate[T constraints.Integer](g func() (T, bool)) Stream[T] {
	return &generateStream[T]{
		g: g,
	}
}
