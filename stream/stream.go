package stream

type Stream[T any] interface {
	Next() (T, bool)
}
