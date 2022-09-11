// Package stream defines a set of operations that provide lazy operations on
// a stream of data values. This package is still experimental, please use at
// your own risk.
package stream

// Stream is an interface that provides a uni-directional, pull-based data
// processing facility.
type Stream[T any] interface {
	// Next retrieves the next item from the stream. It returns (value, at_end).
	// If the stream has no more data, the bool parameter of the return will be
	// true, otherwise it will be false.
	Next() (T, bool)
}
