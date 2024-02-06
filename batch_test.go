package batch_test

import (
	"reflect"
	"testing"

	"github.com/veggiemonk/batch"
)

func TestBatchSlice(t *testing.T) {
	type args struct {
		a []int
		b int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "batch slice 10/3",
			args: args{
				a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				b: 3,
			},
			want: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9, 10},
			},
		},
		{
			name: "batch slice 6/4",
			args: args{
				a: []int{1, 2, 3, 4, 5, 6},
				b: 4,
			},
			want: [][]int{
				{1},
				{2, 3},
				{4},
				{5, 6},
			},
		},
		{
			name: "batch slice 3/3",
			args: args{
				a: []int{1, 2, 3},
				b: 3,
			},
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name: "batch slice 2/3",
			args: args{
				a: []int{1, 2},
				b: 3,
			},
			want: [][]int{{}, {1}, {2}},
		},
		{
			name: "batch slice 2/10",
			args: args{
				a: []int{1, 2},
				b: 10,
			},
			want: [][]int{{}, {}, {}, {}, {1}, {}, {}, {}, {}, {2}},
		},
		{
			name: "batch slice 0/3",
			args: args{
				a: []int{},
				b: 3,
			},
			want: [][]int{{}, {}, {}},
		},
		{
			name: "batch slice 3/0",
			args: args{
				a: []int{1, 2, 3},
				b: 0,
			},
			want: [][]int{},
		},
		{
			name: "batch slice 3/-1",
			args: args{
				a: []int{1, 2, 3},
				b: -1,
			},
			want: [][]int{},
		},
		{
			name: "batch slice 42/10",
			args: args{
				a: genArrayInt(42),
				b: 10,
			},
			want: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
				{17, 18, 19, 20, 21},
				{22, 23, 24, 25},
				{26, 27, 28, 29},
				{30, 31, 32, 33},
				{34, 35, 36, 37},
				{38, 39, 40, 41, 42},
			},
		},
		{
			name: "batch slice 42/7",
			args: args{
				a: genArrayInt(42),
				b: 7,
			},
			want: [][]int{
				{1, 2, 3, 4, 5, 6},
				{7, 8, 9, 10, 11, 12},
				{13, 14, 15, 16, 17, 18},
				{19, 20, 21, 22, 23, 24},
				{25, 26, 27, 28, 29, 30},
				{31, 32, 33, 34, 35, 36},
				{37, 38, 39, 40, 41, 42},
			},
		},
		{
			name: "batch slice 41/7",
			args: args{
				a: genArrayInt(41),
				b: 7,
			},
			want: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10, 11},
				{12, 13, 14, 15, 16, 17},
				{18, 19, 20, 21, 22, 23},
				{24, 25, 26, 27, 28, 29},
				{30, 31, 32, 33, 34, 35},
				{36, 37, 38, 39, 40, 41},
			},
		},
		{
			name: "batch slice 41/5",
			args: args{
				a: genArrayInt(41),
				b: 5,
			},
			want: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8},
				{9, 10, 11, 12, 13, 14, 15, 16},
				{17, 18, 19, 20, 21, 22, 23, 24},
				{25, 26, 27, 28, 29, 30, 31, 32},
				{33, 34, 35, 36, 37, 38, 39, 40, 41},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := batch.Slice(tt.args.a, tt.args.b)
			if len(got) != len(tt.want) {
				t.Errorf("\nBatchSlice() [len = %d]\n        want [len = %d]", len(got), len(tt.want))
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("\nBatchSlice() = %v\n        want = %v", got, tt.want)
			}
		},
		)
	}
}

func genArrayInt(n int) []int {
	var a []int
	for i := 1; i <= n; i++ {
		a = append(a, i)
	}
	return a
}
