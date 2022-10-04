package stream

import "github.com/nadiasvertex/algorithms/cnt"

// Stream is an interface that provides a uni-directional, pull-based data
// processing facility.
type Stream[T any] interface {
	// Next retrieves the next item from the Stream. It returns an Optional[T].
	// If the result is empty it means that the Stream has finished.
	Next() *cnt.Optional[T]
}
