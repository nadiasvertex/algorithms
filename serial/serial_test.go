package serial

import (
	"github.com/nadiasvertex/algorithms/common"
	"reflect"
	"testing"
)

func TestAllOf(t *testing.T) {
	type args struct {
		collection []int
		value      int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "all ones want one",
			args: args{
				collection: []int{1, 1, 1, 1, 1},
				value:      1,
			},
			want: true,
		},
		{
			name: "all ones want 0",
			args: args{
				collection: []int{1, 1, 1, 1, 1},
				value:      0,
			},
			want: false,
		},
		{
			name: "monotonic numbers want 1",
			args: args{
				collection: []int{1, 2, 3, 4, 5},
				value:      1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllOf(tt.args.collection, tt.args.value); got != tt.want {
				t.Errorf("AllOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestAllOfIf(t *testing.T) {
	type args struct {
		collection []int
		pred       common.Predicate[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "all greater than 0 want true",
			args: args{
				collection: []int{1, 2, 3, 4, 5},
				pred: func(value int) bool {
					return value > 0
				},
			}, want: true},
		{name: "all greater than 0 want false",
			args: args{
				collection: []int{-1, -2, -3, -4, -5},
				pred: func(value int) bool {
					return value > 0
				},
			}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllOfIf(tt.args.collection, tt.args.pred); got != tt.want {
				t.Errorf("AllOfIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyOf(t *testing.T) {
	type args struct {
		collection []int
		value      int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "any of all ones is one",
			args: args{
				collection: []int{1, 1, 1, 1, 1},
				value:      1,
			},
			want: true,
		},
		{
			name: "any of all ones is zero",
			args: args{
				collection: []int{1, 1, 1, 1, 1},
				value:      0,
			},
			want: false,
		},
		{
			name: "any of monotonic numbers is one",
			args: args{
				collection: []int{1, 2, 3, 4, 5},
				value:      1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyOf(tt.args.collection, tt.args.value); got != tt.want {
				t.Errorf("AnyOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyOfIf(t *testing.T) {
	type args struct {
		collection []int
		pred       common.Predicate[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "any greater than 0 is true",
			args: args{
				collection: []int{1, 2, 3, 4, 5},
				pred: func(value int) bool {
					return value > 0
				},
			}, want: true},
		{name: "any greater than 0 is false",
			args: args{
				collection: []int{-1, -2, -3, -4, -5},
				pred: func(value int) bool {
					return value > 0
				},
			}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyOfIf(tt.args.collection, tt.args.pred); got != tt.want {
				t.Errorf("AnyOfIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	type args struct {
		x []int
		y []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "compare zeros < ones",
			args: args{
				x: []int{0, 0, 0, 0, 0},
				y: []int{1, 1, 1, 1, 1},
			},
			want: -1,
		},
		{
			name: "compare ones > zeros",
			args: args{
				x: []int{1, 1, 1, 1, 1},
				y: []int{0, 0, 0, 0, 0},
			},
			want: 1,
		},
		{
			name: "compare ones == ones",
			args: args{
				x: []int{1, 1, 1, 1, 1},
				y: []int{1, 1, 1, 1, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		collection []int
		value      int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "collection contains 1",
			args: args{
				collection: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
				value:      1,
			},
			want: true,
		},
		{
			name: "collection does not contain 0",
			args: args{
				collection: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
				value:      0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.collection, tt.args.value); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsIf(t *testing.T) {
	type args struct {
		collection []int
		pred       common.Predicate[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "contains an even number",
			args: args{
				collection: []int{1, 2, 4, 5, 8},
				pred: func(value int) bool {
					return value%2 == 0
				},
			},
			want: true,
		},
		{
			name: "contains no even number",
			args: args{
				collection: []int{1, 3, 5, 5, 7},
				pred: func(value int) bool {
					return value%2 == 0
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsIf(tt.args.collection, tt.args.pred); got != tt.want {
				t.Errorf("ContainsIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	type args struct {
		srcCollection []int
		dstCollection []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "copy from a to b",
			args: args{
				srcCollection: []int{1, 2, 3, 4, 5},
				dstCollection: make([]int, 5),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Copy(tt.args.srcCollection, tt.args.dstCollection)
			if !reflect.DeepEqual(tt.args.srcCollection, tt.args.dstCollection) {
				t.Errorf("Copy() %v != %v", tt.args.srcCollection, tt.args.dstCollection)
			}
		})
	}
}

func TestCopyAppendIf(t *testing.T) {
	type args struct {
		srcCollection []int
		pred          common.Predicate[int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "copy only even numbers",
			args: args{
				srcCollection: []int{1, 2, 3, 4, 5, 6, 7, 8},
				pred: func(value int) bool {
					return value%2 == 0
				},
			},
			want: []int{2, 4, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CopyAppendIf(tt.args.srcCollection, tt.args.pred); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopyAppendIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCopyIf(t *testing.T) {
	type args struct {
		srcCollection []int
		dstCollection []int
		pred          common.Predicate[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "copy even numbers",
			args: args{
				srcCollection: []int{1, 2, 3, 4, 5, 6, 7, 8},
				dstCollection: make([]int, 4),
				pred: func(value int) bool {
					return value%2 == 0
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CopyIf(tt.args.srcCollection, tt.args.dstCollection, tt.args.pred); got != tt.want {
				t.Errorf("CopyIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	type args struct {
		collection []int
		value      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "count ones",
			args: args{
				collection: []int{1, 2, 3, 1, 4, 5, 6, 1, 7, 8, 9, 1},
				value:      1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.args.collection, tt.args.value); got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountIf(t *testing.T) {
	type args struct {
		collection []int
		pred       common.Predicate[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "count even numbers",
			args: args{
				collection: []int{1, 2, 3, 4, 5, 6, 7, 8},
				pred: func(value int) bool {
					return value%2 == 0
				},
			}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountIf(tt.args.collection, tt.args.pred); got != tt.want {
				t.Errorf("CountIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFill(t *testing.T) {
	type args struct {
		collection []int
		value      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "fill a collection with ones",
			args: args{
				collection: make([]int, 5),
				value:      1,
			},
			want: []int{1, 1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fill(tt.args.collection, tt.args.value)
			for i := range tt.args.collection {
				got := tt.args.collection[i]
				if !reflect.DeepEqual(got, tt.args.value) {
					t.Errorf("Fill() = %v, want %v", got, tt.args.value)
				}
			}
		})
	}
}

func TestFind(t *testing.T) {
	type args struct {
		collection []int
		value      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find index of 3",
			args: args{
				collection: []int{1, 2, 3, 4, 5, 6},
				value:      3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Find(tt.args.collection, tt.args.value); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindIf(t *testing.T) {
	type args struct {
		collection []int
		pred       common.Predicate[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find first even number",
			args: args{
				collection: []int{1, 2, 3, 4, 5, 6, 7, 8},
				pred: func(value int) bool {
					return value%2 == 0
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindIf(tt.args.collection, tt.args.pred); got != tt.want {
				t.Errorf("FindIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindIfNot(t *testing.T) {
	type args struct {
		collection []int
		pred       common.Predicate[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find first not even number",
			args: args{
				collection: []int{1, 2, 3, 4, 5, 6, 7, 8},
				pred: func(value int) bool {
					return value%2 == 0
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindIfNot(tt.args.collection, tt.args.pred); got != tt.want {
				t.Errorf("FindIfNot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	type args struct {
		collection []int
		g          common.Generator[int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "generate monotonic int",
			args: args{
				collection: make([]int, 5),
				g: func(index int) int {
					return index
				},
			},
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Generate(tt.args.collection, tt.args.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestHeapSort(t *testing.T) {
	type args struct {
		collection []T
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.collection)
		})
	}
}

func TestInsertionSort(t *testing.T) {
	type args struct {
		collection []T
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.args.collection)
		})
	}
}

func TestIntroSort(t *testing.T) {
	type args struct {
		collection []T
		maxDepth   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IntroSort(tt.args.collection, tt.args.maxDepth)
		})
	}
}

func TestMakeHeap(t *testing.T) {
	type args struct {
		collection []T
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MakeHeap(tt.args.collection)
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		v1 T
		v2 T
	}
	tests := []struct {
		name string
		args args
		want T
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.v1, tt.args.v2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMismatch(t *testing.T) {
	type args struct {
		collection1 []T
		collection2 []T
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mismatch(tt.args.collection1, tt.args.collection2); got != tt.want {
				t.Errorf("Mismatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/

func TestNoneOf(t *testing.T) {
	type args struct {
		collection []int
		value      int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "all ones want none",
			args: args{
				collection: []int{1, 1, 1, 1, 1},
				value:      1,
			},
			want: false,
		},
		{
			name: "all ones want 0",
			args: args{
				collection: []int{1, 1, 1, 1, 1},
				value:      0,
			},
			want: true,
		},
		{
			name: "monotonic numbers want no ones",
			args: args{
				collection: []int{1, 2, 3, 4, 5},
				value:      1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoneOf(tt.args.collection, tt.args.value); got != tt.want {
				t.Errorf("NoneOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoneOfIf(t *testing.T) {
	type args struct {
		collection []int
		pred       common.Predicate[int]
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "all greater than 0 want false",
			args: args{
				collection: []int{1, 2, 3, 4, 5},
				pred: func(value int) bool {
					return value > 0
				},
			}, want: false},
		{name: "all greater than 0 want true",
			args: args{
				collection: []int{-1, -2, -3, -4, -5},
				pred: func(value int) bool {
					return value > 0
				},
			}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NoneOfIf(tt.args.collection, tt.args.pred); got != tt.want {
				t.Errorf("NoneOfIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartition(t *testing.T) {
	type args struct {
		collection []int
		pred       common.Predicate[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "partition even and odd numbers",
			args: args{
				collection: []int{1, 2, 3, 4, 5, 6, 7, 8},
				pred: func(value int) bool {
					return value%2 == 0
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Partition(tt.args.collection, tt.args.pred); got != tt.want {
				t.Errorf("Partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestRandomShuffle(t *testing.T) {
	type args struct {
		collection []T
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RandomShuffle(tt.args.collection)
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		collection []T
		value      T
	}
	tests := []struct {
		name string
		args args
		want []T
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.args.collection, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveIf(t *testing.T) {
	type args struct {
		collection []T
		pred       Predicate
	}
	tests := []struct {
		name string
		args args
		want []T
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveIf(tt.args.collection, tt.args.pred); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveIf() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/

func TestReplace(t *testing.T) {
	type args struct {
		collection []int
		oldValue   int
		newValue   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "replace 1 with 0",
			args: args{
				collection: []int{0, 1, 2, 3, 1, 4, 5, 6, 1, 8, 9},
				oldValue:   1,
				newValue:   0,
			},
			want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Replace(tt.args.collection, tt.args.oldValue, tt.args.newValue); got != tt.want {
				t.Errorf("Replace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceIf(t *testing.T) {
	type args struct {
		collection []int
		pred       common.Predicate[int]
		newValue   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "replace even numbers with 1",
			args: args{
				collection: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				pred: func(value int) bool {
					return value%2 == 0
				},
				newValue: 1,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceIf(tt.args.collection, tt.args.pred, tt.args.newValue); got != tt.want {
				t.Errorf("ReplaceIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		collection []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "reverse monotonic",
			args: args{
				collection: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Reverse(tt.args.collection)
			if !reflect.DeepEqual(tt.args.collection, tt.want) {
				t.Errorf("Reverse() = %v, want %v", tt.args.collection, tt.want)
			}
		})
	}
}

func TestReverseCopy(t *testing.T) {
	type args struct {
		collection []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "reverse monotonic",
			args: args{
				collection: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			},
			want: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseCopy(tt.args.collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	type args struct {
		collection []int
		nFirst     int
	}
	tests := []struct {
		name          string
		args          args
		want          int
		newCollection []int
	}{
		{
			name: "rotate entire collection",
			args: args{
				collection: []int{1, 2, 3, 4, 5, 6, 7, 8},
				nFirst:     5,
			},
			want:          3,
			newCollection: []int{6, 7, 8, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rotate(tt.args.collection, tt.args.nFirst); got != tt.want {
				t.Errorf("Rotate() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.args.collection, tt.newCollection) {
				t.Errorf("Rotate() = %v, want %v", tt.args.collection, tt.newCollection)
			}
		})
	}
}

func TestRotateRange(t *testing.T) {
	type args struct {
		collection []int
		first      int
		nFirst     int
		last       int
	}
	tests := []struct {
		name          string
		args          args
		want          int
		newCollection []int
	}{
		{
			name: "rotate inner range",
			args: args{
				collection: []int{1, 2, 3, 4, 5, 6, 7, 8},
				first:      2,
				nFirst:     4,
				last:       6,
			},
			want:          4,
			newCollection: []int{1, 2, 5, 6, 3, 4, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotateRange(tt.args.collection, tt.args.first, tt.args.nFirst, tt.args.last); got != tt.want {
				t.Errorf("RotateRange() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.args.collection, tt.newCollection) {
				t.Errorf("RotateRange() = %v, want %v", tt.args.collection, tt.newCollection)
			}
		})
	}
}

/*
	func TestSort(t *testing.T) {
		type args struct {
			collection []T
		}
		tests := []struct {
			name string
			args args
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				Sort(tt.args.collection)
			})
		}
	}

	func TestSortPartition(t *testing.T) {
		type args struct {
			collection []T
			lo         int
			hi         int
		}
		tests := []struct {
			name string
			args args
			want int
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := SortPartition(tt.args.collection, tt.args.lo, tt.args.hi); got != tt.want {
					t.Errorf("SortPartition() = %v, want %v", got, tt.want)
				}
			})
		}
	}
*/

func TestSwap(t *testing.T) {
	v1 := 5
	v2 := 10

	type args struct {
		i1 *int
		i2 *int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "swap ints",
			args: args{
				&v1,
				&v2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Swap(tt.args.i1, tt.args.i2)
			if v1 != 10 || v2 != 5 {
				t.Errorf("Swap() = (%d, %d), want (10, 5)", v1, v2)
			}
		})
	}
}

func TestSwapIndex(t *testing.T) {
	type args struct {
		collection []int
		i1         int
		i2         int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "swap first two items",
			args: args{
				collection: []int{0, 1, 2, 3, 4, 5},
				i1:         0,
				i2:         1,
			},
			want: []int{1, 0, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SwapIndex(tt.args.collection, tt.args.i1, tt.args.i2)
			if !reflect.DeepEqual(tt.args.collection, tt.want) {
				t.Errorf("SwapIndex() = %v, want %v", tt.args.collection, tt.want)
			}
		})
	}
}

func TestTransform(t *testing.T) {
	type args struct {
		srcCollection []int
		dstCollection []int
		unary         common.UnaryTransform[int, int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "double ints",
			args: args{
				srcCollection: []int{1, 2, 3, 4, 5, 6},
				dstCollection: make([]int, 6),
				unary: func(value int) int {
					return value * 2
				},
			},
			want: []int{2, 4, 6, 8, 10, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Transform(tt.args.srcCollection, tt.args.dstCollection, tt.args.unary); !reflect.DeepEqual(got,
				tt.want) {
				t.Errorf("Transform() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransformAppend(t *testing.T) {
	type args struct {
		srcCollection []int
		unary         common.UnaryTransform[int, int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "double ints",
			args: args{
				srcCollection: []int{1, 2, 3, 4, 5, 6},
				unary: func(value int) int {
					return value * 2
				},
			},
			want: []int{2, 4, 6, 8, 10, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransformAppend(tt.args.srcCollection, tt.args.unary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransformAppend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransformAppendIf(t *testing.T) {
	type args struct {
		srcCollection []int
		pred          common.Predicate[int]
		unary         common.UnaryTransform[int, int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "double ints",
			args: args{
				srcCollection: []int{1, 2, 3, 4, 5, 6},
				pred: func(value int) bool {
					return value%2 == 0
				},
				unary: func(value int) int {
					return value * 2
				},
			},
			want: []int{4, 8, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransformAppendIf(tt.args.srcCollection, tt.args.pred, tt.args.unary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransformAppendIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransformCopyIf(t *testing.T) {
	type args struct {
		srcCollection []int
		dstCollection []int
		pred          common.Predicate[int]
		unary         common.UnaryTransform[int, int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "double even ints",
			args: args{
				srcCollection: []int{1, 2, 3, 4, 5, 6},
				dstCollection: make([]int, 6),
				pred: func(value int) bool {
					return value%2 == 0
				},
				unary: func(value int) int {
					return value * 2
				},
			},
			want: []int{1, 4, 3, 8, 5, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransformCopyIf(tt.args.srcCollection, tt.args.dstCollection, tt.args.pred,
				tt.args.unary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransformCopyIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransformIf(t *testing.T) {
	type args struct {
		srcCollection []int
		dstCollection []int
		pred          common.Predicate[int]
		unary         common.UnaryTransform[int, int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "double even ints",
			args: args{
				srcCollection: []int{1, 2, 3, 4, 5, 6},
				dstCollection: make([]int, 6),
				pred: func(value int) bool {
					return value%2 == 0
				},
				unary: func(value int) int {
					return value * 2
				},
			},
			want: []int{4, 8, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TransformIf(tt.args.srcCollection, tt.args.dstCollection, tt.args.pred,
				tt.args.unary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransformIf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnique(t *testing.T) {
	type args struct {
		collection []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "remove duplicate values",
			args: args{
				collection: []int{1, 1, 1, 1, 1, 2, 3, 4, 4, 4, 4, 4, 5, 6, 7, 8},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "don't die on small collection",
			args: args{
				collection: []int{1},
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
