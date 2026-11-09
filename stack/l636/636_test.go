package l636

import (
	"slices"
	"testing"
)

func TestExclusiveTime(t *testing.T) {
	tests := []struct {
		name string
		n    int
		logs []string
		want []int
	}{
		// LeetCode examples
		{
			"example 1", 2,
			[]string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"},
			[]int{3, 4},
		},
		{
			"example 2 recursion", 1,
			[]string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"},
			[]int{8},
		},
		{
			"example 3", 2,
			[]string{"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"},
			[]int{7, 1},
		},

		// single call
		{
			"single unit call", 1,
			[]string{"0:start:0", "0:end:0"},
			[]int{1},
		},
		{
			"start not at zero", 1,
			[]string{"0:start:5", "0:end:9"},
			[]int{5},
		},
		{
			"large timestamps", 1,
			[]string{"0:start:1000000000", "0:end:1000000000"},
			[]int{1},
		},

		// multiple functions
		{
			"unused functions get zero", 3,
			[]string{"0:start:0", "0:end:4"},
			[]int{5, 0, 0},
		},
		{
			"sequential calls", 2,
			[]string{"0:start:0", "0:end:2", "1:start:3", "1:end:4"},
			[]int{3, 2},
		},
		{
			"multi digit function id", 11,
			[]string{"10:start:0", "10:end:9"},
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10},
		},

		// nesting
		{
			"single unit child", 2,
			[]string{"0:start:0", "1:start:1", "1:end:1", "0:end:2"},
			[]int{2, 1},
		},
		{
			"deep nesting", 3,
			[]string{"0:start:0", "1:start:1", "2:start:2", "2:end:3", "1:end:4", "0:end:5"},
			[]int{2, 2, 2},
		},
		{
			"back to back children", 3,
			[]string{"0:start:0", "1:start:1", "1:end:2", "2:start:3", "2:end:3", "0:end:5"},
			[]int{3, 2, 1},
		},
		{
			"deep recursion same function", 1,
			[]string{"0:start:0", "0:start:1", "0:start:2", "0:end:3", "0:end:4", "0:end:5"},
			[]int{6},
		},
		{
			"child takes all parent time", 2,
			[]string{"0:start:0", "1:start:1", "1:end:8", "0:end:9"},
			[]int{2, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exclusiveTime(tt.n, tt.logs); !slices.Equal(got, tt.want) {
				t.Errorf("exclusiveTime(%d, %q) = %v, want %v", tt.n, tt.logs, got, tt.want)
			}
		})
	}
}
