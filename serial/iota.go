package serial

import "golang.org/x/exp/constraints"

// Iota fills collection with monotonically increasing integers
// starting from 0.
func Iota[T constraints.Integer](collection []T) []T {
	return IotaWithStart(collection, 0)
}

// IotaN fills a collection with N monotonically increasing integers
// starting from 0.
func IotaN[T constraints.Integer](count int) []T {
	return Iota(make([]T, count))
}

// IotaNWithStart fills a collection with N monotonically increasing integers
// starting from 0.
func IotaNWithStart[T constraints.Integer](count int, start T) []T {
	return IotaWithStart(make([]T, count), start)
}

// IotaWithStart fills a collection with monotonically increasing
// integers starting at start.
func IotaWithStart[T constraints.Integer](collection []T, start T) []T {
	value := start
	for i := range collection {
		collection[i] = value
		value++
	}

	return collection
}
