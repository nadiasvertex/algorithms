package stream

import (
	"github.com/nadiasvertex/algorithms/cnt"
	"golang.org/x/exp/constraints"
)

type generateStream[T any] struct {
	g func() *cnt.Optional[T]
}

func (s *generateStream[T]) Next() *cnt.Optional[T] {
	return s.g()
}

// Generate creates a stream that generates user-defined values. When the
// generator wants to terminate the stream, it should return true as it's
// second resut argument.
func Generate[T constraints.Integer](g func() *cnt.Optional[T]) Stream[T] {
	return &generateStream[T]{
		g: g,
	}
}
