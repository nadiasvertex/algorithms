package parallel

import (
	"github.com/nadiasvertex/algorithms/common"
	"reflect"
	"testing"

	"github.com/nadiasvertex/algorithms/serial"
)

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
				collection: Fill(make([]int, 5000), 1),
				value:      1,
			},
			want: 5000,
		},
		{
			name: "count ones get zero",
			args: args{
				collection: Fill(make([]int, 5000), 2),
				value:      1,
			},
			want: 0,
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
			name: "find index of 4000",
			args: args{
				collection: serial.Generate(make([]int, 5000), func(index int) int {
					return index
				}),
				value: 4000,
			},
			want: 4000,
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
				collection: serial.Generate(make([]int, 5000), func(index int) int {
					if index%2 == 0 {
						return 1
					} else {
						return index
					}
				}),
				oldValue: 1,
				newValue: 0,
			},
			want: 2501},
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
				collection: serial.Generate(make([]int, 5000), func(index int) int {
					return index
				}),
				pred: func(value int) bool {
					return value%2 == 0
				},
				newValue: 1,
			},
			want: 2500,
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
