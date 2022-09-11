package stream

import "golang.org/x/exp/constraints"

type iotaStream[T constraints.Integer] struct {
	next T
}

func (s *iotaStream[T]) Next() (T, bool) {
	current := s.next
	s.next++
	return current, false
}

// Iota creates a stream that provides an infinite number of monotonically
// increasing integers.
func Iota[T constraints.Integer]() Stream[T] {
	return &iotaStream[T]{
		next: 0,
	}
}

// IotaWithStart creates a stream just like Iota, but allows you to set the
// starting point.
func IotaWithStart[T constraints.Integer](start T) Stream[T] {
	return &iotaStream[T]{
		next: start,
	}
}
