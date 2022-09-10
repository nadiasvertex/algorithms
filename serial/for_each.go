package serial

import "github.com/nadiasvertex/algorithms/common"

// ForEach runs op on each item in the element
func ForEach[T any](collection []T, op common.UnaryOp[T]) {
	for _, item := range collection {
		op(item)
	}
}
