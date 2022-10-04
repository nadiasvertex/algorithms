package stream

import "github.com/nadiasvertex/algorithms/cnt"

// OneOrNone reads one item from the stream and returns it or an
// empty optional.
func OneOrNone[T any](input Stream[T]) *cnt.Optional[T] {
	return input.Next()
}
