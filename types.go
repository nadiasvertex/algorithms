package algorithms

type Predicate[T any] func(value T) bool

type UnaryTransform[T1, T2 any] func(value T1) T2

type Generator[T any] func() T
