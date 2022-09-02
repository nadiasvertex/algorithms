package serial

import "github.com/nadiasvertex/algorithms/common"

// Transform transforms all items in srcCollection to items in dstCollection. dstCollection must be at
// least as big as srcCollection.
func Transform[T1, T2 any](srcCollection []T1, dstCollection []T2, unary common.UnaryTransform[T1, T2]) {
	for i, item := range srcCollection {
		dstCollection[i] = unary(item)
	}
}

// TransformAppend transforms all items in srcCollection to items in dstCollection. dstCollection must be at
// least as big as srcCollection.
func TransformAppend[T1, T2 any](srcCollection []T1, unary common.UnaryTransform[T1, T2]) []T2 {
	var dstCollection []T2
	for _, item := range srcCollection {
		dstCollection = append(dstCollection, unary(item))
	}
	return dstCollection
}

// TransformIf transforms all items in srcCollection that match the given predicate. dstCollection must
// be at least as large as srcCollection. The return value indicates how many items werre copied.
func TransformIf[T1, T2 any](srcCollection []T1, dstCollection []T2, pred common.Predicate[T1], unary common.UnaryTransform[T1,
	T2]) int {
	output := 0
	for _, item := range srcCollection {
		if pred(item) {
			dstCollection[output] = unary(item)
			output++
		}
	}

	return output
}

// TransformAppendIf transforms all items in srcCollection that match the given predicate.
// Transformed items are appended to a new collection. The return value is the
// collection of all transformed items.
func TransformAppendIf[T1, T2 any](srcCollection []T1, pred common.Predicate[T1], unary common.UnaryTransform[T1,
	T2]) []T2 {
	var dstCollection []T2
	for _, item := range srcCollection {
		if pred(item) {
			dstCollection = append(dstCollection, unary(item))
		}
	}
	return dstCollection
}

// TransformCopyIf transforms all items in srcCollection that match the given predicate.
// Items which don't match are just copied. dstCollection must
// be at least as large as srcCollection. The return value indicates how many items were transformed.
func TransformCopyIf[T any](srcCollection []T, dstCollection []T, pred common.Predicate[T], unary common.UnaryTransform[T,
	T]) {
	for i, item := range srcCollection {
		if pred(item) {
			dstCollection[i] = unary(item)
		} else {
			dstCollection[i] = item
		}
	}
}
