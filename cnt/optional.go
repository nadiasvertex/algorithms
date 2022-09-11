package cnt

import "github.com/nadiasvertex/algorithms/common"

// Optional encapsulates a value that may or may not be present.
type Optional[T any] interface {
	// Value provides the contained value, or panics if there is none.
	Value() T

	// ValueOr provides the contained value, or if there is none it provides
	// defaultValue.
	ValueOr(defaultValue T) T

	// AndThen extracts the contained value if present, runs the transform
	// function on it, and returns the new value. If not present it returns
	// an empty value of type T.
	AndThen(transform common.UnaryTransform[T, T]) T

	// Transform extracts the contained value if present, runs the transform
	// function on it, and returns an Optional containing the new value. If
	// not present it returns an empty optional.
	Transform(transform common.UnaryTransform[T, T]) Optional[T]

	// OrElse returns this object if it has a value, or executes hte cons
	// function and wraps the result in a new optional value.
	OrElse(cons common.UnaryGenerator[T]) Optional[T]

	// HasValue returns true if there is an contained value, or false otherwise.
	HasValue() bool
}

// fullOptional implements an Optional type that contains a value.
type fullOptional[T any] struct {
	value T
}

func (f *fullOptional[T]) Value() T {
	return f.value
}

func (f *fullOptional[T]) ValueOr(T) T {
	return f.value
}

func (f *fullOptional[T]) AndThen(xform common.UnaryTransform[T, T]) T {
	return xform(f.value)
}

func (f *fullOptional[T]) Transform(transform common.UnaryTransform[T, T]) Optional[T] {
	return &fullOptional[T]{value: transform(f.value)}
}

func (f *fullOptional[T]) OrElse(cons common.UnaryGenerator[T]) Optional[T] {
	return f
}

func (f *fullOptional[T]) HasValue() bool {
	return true
}

// emptyOptional implements a type that has no value.
type emptyOptional[T any] struct{}

func (e *emptyOptional[T]) Value() T {
	panic("Optional has no value.")
}

func (e *emptyOptional[T]) ValueOr(t T) T {
	return t
}

func (e *emptyOptional[T]) AndThen(xform common.UnaryTransform[T, T]) T {
	var output T
	return output
}

func (e *emptyOptional[T]) Transform(xform common.UnaryTransform[T, T]) Optional[T] {
	return e
}

func (e *emptyOptional[T]) OrElse(cons common.UnaryGenerator[T]) Optional[T] {
	return &fullOptional[T]{value: cons()}
}

func (e *emptyOptional[T]) HasValue() bool {
	return false
}

// AndThen returns the result of invocation of transform on the contained value if it
// exists. Otherwise, it returns an empty value of the return type.
func AndThen[T1, T2 any](in Optional[T1], transform common.UnaryTransform[T1, T2]) T2 {
	if !in.HasValue() {
		var v T2
		return v
	}

	return transform(in.Value())
}

// Transform returns an [Optional] that contains the result of invocation of transform
// on the contained value if it contains a value. Otherwise, returns an empty Optional
// of such type.
func Transform[T1, T2 any](in Optional[T1], transform common.UnaryTransform[T1, T2]) Optional[T2] {
	if !in.HasValue() {
		return &emptyOptional[T2]{}
	}

	return &fullOptional[T2]{value: transform(in.Value())}
}

// NullOpt returns an empty [Optional] of T.
func NullOpt[T any]() Optional[T] {
	return &emptyOptional[T]{}
}

// MakeOptional wraps the given value in a new optional object.
func MakeOptional[T any](value T) Optional[T] {
	return &fullOptional[T]{value: value}
}
