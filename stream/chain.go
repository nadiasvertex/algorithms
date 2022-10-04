package stream

import (
	"github.com/nadiasvertex/algorithms/cnt"
	"github.com/nadiasvertex/algorithms/common"
	"golang.org/x/exp/constraints"
)

type Chain[T any] struct {
	Stream Stream[T]
}

func (s Chain[T]) Filter(pred common.Predicate[T]) Chain[T] {
	return Chain[T]{Stream: Filter(s.Stream, pred)}
}

func (s Chain[T]) Take(limit int) Chain[T] {
	return Chain[T]{Stream: Take(s.Stream, limit)}
}

func (s Chain[T]) Skip(limit int) Chain[T] {
	return Chain[T]{Stream: Skip(s.Stream, limit)}
}

func (s Chain[T]) MergeSort(o Chain[T], pred common.BinaryPredicate[T]) Chain[T] {
	return Chain[T]{Stream: MergeSort[T](s.Stream, o.Stream, pred)}
}

func (s Chain[T]) ForEach(pred common.UnaryOp[T]) {
	ForEach(s.Stream, pred)
}

func (s Chain[T]) ToSlice() []T {
	return ToSlice(s.Stream)
}

func (s Chain[T]) OneOrNone() *cnt.Optional[T] {
	return OneOrNone(s.Stream)
}

func Pipe[T any](s Stream[T]) Chain[T] {
	return Chain[T]{Stream: s}
}

func Slice[T any](data []T) Chain[T] {
	return Chain[T]{Stream: FromSlice(data)}
}

func Generator[T any](g func() *cnt.Optional[T]) Chain[T] {
	return Chain[T]{Stream: Generate(g)}
}

func Monotonic[T constraints.Integer]() Chain[T] {
	return Chain[T]{Stream: Iota[T]()}
}

func MonotonicWithStart[T constraints.Integer](start T) Chain[T] {
	return Chain[T]{Stream: IotaWithStart(start)}
}
