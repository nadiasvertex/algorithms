package stream

import "github.com/nadiasvertex/algorithms/common"

func ForEach[T any](s Stream[T], op common.UnaryOp[T]) {
	for {
		if v, atEnd := s.Next(); atEnd {
			return
		} else {
			op(v)
		}
	}
}
