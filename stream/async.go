package stream

import (
	"sync/atomic"
)

type asyncStream[T any] struct {
	input   Stream[T]
	channel chan T
	closed  atomic.Bool
}

func (s *asyncStream[T]) Next() (T, bool) {
	v, ok := <-s.channel
	return v, !ok
}

func (s *asyncStream[T]) processAsync() {
	for {
		v, atEnd := s.input.Next()
		if atEnd {
			s.closed.Store(true)
			close(s.channel)
			return
		}

		s.channel <- v
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
