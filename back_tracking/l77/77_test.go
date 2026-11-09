package l77

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		name string
		n, k int
		want [][]int
	}{
		{
			name: "example one",
			n:    4,
			k:    2,
			want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name: "example two",
			n:    1,
			k:    1,
			want: [][]int{{1}},
		},
		{
			name: "single elements",
			n:    4,
			k:    1,
			want: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name: "all elements",
			n:    4,
			k:    4,
			want: [][]int{{1, 2, 3, 4}},
		},
		{
			name: "three element combinations",
			n:    5,
			k:    3,
			want: [][]int{{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5}, {1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5}},
		},
		{
			name: "zero elements",
			n:    3,
			k:    0,
			want: [][]int{{}},
		},
		{
			name: "too many elements",
			n:    2,
			k:    3,
			want: [][]int{},
		},
		{
			name: "negative size",
			n:    2,
			k:    -1,
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combine(tt.n, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combine(%d, %d) = %v, want %v", tt.n, tt.k, got, tt.want)
			}
		})
	}
}

func TestCombineIndependentResults(t *testing.T) {
	got := combine(4, 2)
	got[0][0] = 99
	want := [][]int{{99, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("mutating one combination changed other combinations: %v", got)
	}
}
