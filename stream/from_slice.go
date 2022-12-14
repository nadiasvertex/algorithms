package stream

import (
	"github.com/nadiasvertex/algorithms/cnt"
)

type fromSliceStream[T any] struct {
	elements []T
	index    int
}

func (s *fromSliceStream[T]) Next() *cnt.Optional[T] {
	if s.index == len(s.elements) {
		return cnt.NullOpt[T]()
	}

	v := s.elements[s.index]
	s.index++
	return cnt.Some(v)
}

// FromSlice creates a Stream that provides all the elements in the slice, one
// at a time. When it reaches the end is terminates the Stream.
func FromSlice[T any](data []T) Stream[T] {
	return &fromSliceStream[T]{
		elements: data,
		index:    0,
	}
}
