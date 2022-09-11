package stream

// ToSlice drains a stream and returns the result as a slice.
func ToSlice[T any](input Stream[T]) []T {
	var output []T
	for {
		_ := input.Next()
		if atEnd {
			return output
		}

		output = append(output, v)
	}
}
