package stream

import (
	"github.com/nadiasvertex/algorithms/cnt"
	"sync/atomic"
)

type asyncStream[T any] struct {
	input   Stream[T]
	channel chan T
	closed  atomic.Bool
}

func (s *asyncStream[T]) Next() cnt.Optional[T] {
	if v, ok := <-s.channel; ok {
		return cnt.MakeOptional(v)
	}
	return cnt.NullOpt[T]()
}

func (s *asyncStream[T]) processAsync() {
	for {
		if v := s.input.Next(); !v.HasValue() {
			s.closed.Store(true)
			close(s.channel)
			return
		} else {
			s.channel <- v.Value()
		}
	}
}

// Async create a new asynchronous stream. The input stream will be read from
// a go routine, and its results transferred over a channel. That means
// whatever work the input stream needs to do can be detached from the work
// happening in this thread.
func Async[T any](input Stream[T]) Stream[T] {
	s := &asyncStream[T]{
		input:   input,
		channel: make(chan T),
	}

	go s.processAsync()
	return s
}
