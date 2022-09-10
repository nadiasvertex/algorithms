package stream

func ToSlice[T any](input Stream[T]) []T {
	var output []T
	for {
		v, atEnd := input.Next()
		if atEnd {
			return output
		}

		output = append(output, v)
	}
}
