package serial

// Swap swaps the values of i1 and i2.
func Swap[T any](i1, i2 *T) {
	i3 := *i1
	*i1 = *i2
	*i2 = i3
}

// SwapIndex swaps the values of the elements in collection at i1 and i2.
func SwapIndex[T any](collection []T, i1, i2 int) {
	i3 := collection[i1]
	collection[i1] = collection[i2]
	collection[i2] = i3
}
