package stream

import (
	"github.com/nadiasvertex/algorithms/cnt"
	"github.com/nadiasvertex/algorithms/common"
)

type transformStream[T1, T2 any] struct {
	input Stream[T1]
	xform common.UnaryTransform[T1, T2]
}

func (s *transformStream[T1, T2]) Next() cnt.Optional[T2] {
	return cnt.Transform(s.input.Next(), s.xform)
}

// Transform adds a stream step that transforms a value into another value, optionally
// of another type.
func Transform[T1, T2 any](input Stream[T1], xform common.UnaryTransform[T1, T2]) Stream[T2] {
	return &transformStream[T1, T2]{
		input: input,
		xform: xform,
	}
}
