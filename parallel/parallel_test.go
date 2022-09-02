package parallel

import (
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
