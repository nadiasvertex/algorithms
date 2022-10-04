package stream

import (
	"github.com/nadiasvertex/algorithms/cnt"
)

type generateStream[T any] struct {
	g func() *cnt.Optional[T]
}

func (s *generateStream[T]) Next() *cnt.Optional[T] {
	return s.g()
}

// Generate creates a Stream that generates user-defined values. When the
// generator wants to terminate the Stream, it should return true as it's
// second resut argument.
func Generate[T any](g func() *cnt.Optional[T]) Stream[T] {
	return &generateStream[T]{
		g: g,
	}
}
