// Package stream defines a set of operations that provide lazy operations on
// a stream of data values. This package is still experimental, please use at
// your own risk.
package stream

import "github.com/nadiasvertex/algorithms/cnt"

// Stream is an interface that provides a uni-directional, pull-based data
// processing facility.
type Stream[T any] interface {
	// Next retrieves the next item from the stream. It returns an Optional[T].
	// If the result is empty it means that the stream has finished.
	Next() *cnt.Optional[T]
}
