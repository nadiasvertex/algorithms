package stream

import "golang.org/x/exp/constraints"

type IotaStream[T constraints.Integer] struct {
	next T
}

func (s *IotaStream[T]) Next() (T, bool) {
	current := s.next
	s.next++
	return current, false
}

func Iota[T constraints.Integer]() Stream[T] {
	return &IotaStream[T]{
		next: 0,
	}
}

func IotaWithStart[T constraints.Integer](start T) Stream[T] {
	return &IotaStream[T]{
		next: start,
	}
}
