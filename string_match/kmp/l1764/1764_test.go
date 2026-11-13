package l1764

import "testing"

func TestCanChoose(t *testing.T) {
	tests := []struct {
		name   string
		groups [][]int
		nums   []int
		want   bool
	}{
		{"example", [][]int{{1, -1, -1}, {3, -2, 0}}, []int{1, -1, 0, 1, -1, -1, 3, -2, 0}, true},
		{"gaps allowed", [][]int{{1, 2}, {5}}, []int{1, 2, 3, 4, 5}, true},
		{"adjacent groups", [][]int{{1, 2}, {3}}, []int{1, 2, 3}, true},
		{"overlap forbidden", [][]int{{1, 2}, {2, 3}}, []int{1, 2, 3, 9}, false},
		{"wrong order", [][]int{{2}, {1}}, []int{1, 2}, false},
		{"not subsequence", [][]int{{1, 3}}, []int{1, 2, 3}, false},
		{"fallback", [][]int{{1, 1, 2}}, []int{1, 1, 1, 2}, true},
		{"last group absent", [][]int{{1}, {4}}, []int{1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canChoose(tt.groups, tt.nums); got != tt.want {
				t.Errorf("canChoose(%v, %v) = %v, want %v", tt.groups, tt.nums, got, tt.want)
			}
		})
	}
}
