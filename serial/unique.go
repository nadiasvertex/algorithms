package serial

import "github.com/nadiasvertex/algorithms/common"

// Unique eliminates all except the first element from every consecutive group
// of equivalent elements from the collection.
func Unique[T comparable](collection []T) []T {
	if len(collection) < 2 {
		return collection
	}

	cursor := 0
	for index := range collection {
		if collection[cursor] != collection[index] {
			cursor++
			if cursor != index {
				collection[cursor] = collection[index]
			}
		}
	}

	return collection[0 : cursor+1]
}

// UniqueIf eliminates all except the first element from every consecutive group
// of equivalent elements from the collection. Elements are compared using a
// binary predicate.
func UniqueIf[T any](collection []T, pred common.BinaryPredicate[T]) []T {
	if len(collection) < 2 {
		return collection
	}

	cursor := 0
	for index := range collection {
		if !pred(collection[cursor], collection[index]) {
			cursor++
			if cursor != index {
				collection[cursor] = collection[index]
			}
		}
	}

	return collection[0 : cursor+1]
}
