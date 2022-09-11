package stream

// ToSlice drains a stream and returns the result as a slice.
func ToSlice[T any](input Stream[T]) []T {
	var output []T
	for {
		if v := input.Next(); !v.HasValue() {
			return output
		} else {
			output = append(output, v.Value())
		}
	}
}
