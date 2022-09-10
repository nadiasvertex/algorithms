package stream

import "github.com/nadiasvertex/algorithms/common"

type TransformStream[T1, T2 any] struct {
	input Stream[T1]
	xform common.UnaryTransform[T1, T2]
}

func (s *TransformStream[T1, T2]) Next() (T2, bool) {
	if v, atEnd := s.input.Next(); atEnd {
		var empty T2
		return empty, atEnd
	} else {
		return s.xform(v), false
	}
}

func Transform[T1, T2 any](input Stream[T1], xform common.UnaryTransform[T1, T2]) Stream[T2] {
	return &TransformStream[T1, T2]{
		input: input,
		xform: xform,
	}
}
