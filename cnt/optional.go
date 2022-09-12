package cnt

import "github.com/nadiasvertex/algorithms/common"

// Optional contains a value that may or may not be present.
type Optional[T any] struct {
	value T
}

// Value provides the contained value, or panics if there is none.
func (opt *Optional[T]) Value() T {
	if opt == nil {
		panic("Optional is empty.")
	}
	return opt.value
}

// ValueOr provides the contained value, or if there is none it provides
// defaultValue.
func (opt *Optional[T]) ValueOr(value T) T {
	if opt != nil {
		return opt.value
	}
	return value
}

// AndThen extracts the contained value if present, runs the transform
// function on it, and returns the new value. If not present it returns
// an empty value of type T.
func (opt *Optional[T]) AndThen(xform common.UnaryTransform[T, T]) T {
	if opt != nil {
		return xform(opt.value)
	}
	var value T
	return value
}

// Transform extracts the contained value if present, runs the transform
// function on it, and returns an Optional containing the new value. If
// not present it returns an empty optional.
func (opt *Optional[T]) Transform(transform common.UnaryTransform[T, T]) *Optional[T] {
	if opt != nil {
		return &Optional[T]{value: transform(opt.value)}
	}
	return nil
}

// OrElse returns this object if it has a value, or executes hte cons
// function and wraps the result in a new optional value.
func (opt *Optional[T]) OrElse(cons common.UnaryGenerator[T]) *Optional[T] {
	if opt != nil {
		return opt
	}
	return MakeOptional(cons())
}

// HasValue returns true if there is an contained value, or false otherwise.
func (opt *Optional[T]) HasValue() bool {
	return opt != nil
}

// AndThen returns the result of invocation of transform on the contained value if it
// exists. Otherwise, it returns an empty value of the return type.
func AndThen[T1, T2 any](in *Optional[T1], transform common.UnaryTransform[T1, T2]) T2 {
	if !in.HasValue() {
		var v T2
		return v
	}

	return transform(in.Value())
}

// Transform returns an [Optional] that contains the result of invocation of transform
// on the contained value if it contains a value. Otherwise, returns an empty Optional
// of such type.
func Transform[T1, T2 any](in *Optional[T1], transform common.UnaryTransform[T1, T2]) *Optional[T2] {
	if !in.HasValue() {
		return nil
	}

	return &Optional[T2]{value: transform(in.Value())}
}

// NullOpt returns an empty [Optional] of T.
func NullOpt[T any]() *Optional[T] {
	return nil
}

// MakeOptional wraps the given value in a new optional object.
func MakeOptional[T any](value T) *Optional[T] {
	return &Optional[T]{value: value}
}
