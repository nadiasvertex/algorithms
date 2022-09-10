package common

type Predicate[T any] func(value T) bool
type BinaryPredicate[T any] func(value1, value2 T) bool

type UnaryTransform[T1, T2 any] func(value T1) T2
type UnaryOp[T any] func(value T)

type Generator[T any] func(index int) T
