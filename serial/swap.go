package serial

func Swap[T any](i1, i2 *T) {
	i3 := *i1
	*i1 = *i2
	*i2 = i3
}

func SwapIndex[T any](collection []T, i1, i2 int) {
	i3 := collection[i1]
	collection[i1] = collection[i2]
	collection[i2] = i3
}
