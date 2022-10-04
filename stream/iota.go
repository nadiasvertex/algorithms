package stream

import (
	"github.com/nadiasvertex/algorithms/cnt"
	"golang.org/x/exp/constraints"
)

type iotaStream[T constraints.Integer] struct {
	next T
}

func (s *iotaStream[T]) Next() *cnt.Optional[T] {
	current := cnt.Some(s.next)
	s.next++
	return current
}

// Iota creates a Stream that provides an infinite number of monotonically
// increasing integers.
func Iota[T constraints.Integer]() Stream[T] {
	return &iotaStream[T]{
		next: 0,
	}
}

// IotaWithStart creates a Stream just like Iota, but allows you to set the
// starting point.
func IotaWithStart[T constraints.Integer](start T) Stream[T] {
	return &iotaStream[T]{
		next: start,
	}
}
