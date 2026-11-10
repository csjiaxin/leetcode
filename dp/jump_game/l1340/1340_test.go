package l1340

import "testing"

func TestMaxJumps(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		d    int
		want int
	}{
		{name: "example one", arr: []int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, d: 2, want: 4},
		{name: "example two", arr: []int{3, 3, 3, 3, 3}, d: 3, want: 1},
		{name: "example three", arr: []int{7, 6, 5, 4, 3, 2, 1}, d: 1, want: 7},
		{name: "single index counts as visited", arr: []int{5}, d: 1, want: 1},
		{name: "two descending values", arr: []int{5, 1}, d: 1, want: 2},
		{name: "two ascending values", arr: []int{1, 5}, d: 1, want: 2},
		{name: "increasing chain moves left", arr: []int{1, 2, 3, 4, 5}, d: 1, want: 5},
		{name: "equal height blocks farther lower value", arr: []int{5, 5, 1}, d: 2, want: 2},
		{name: "taller height blocks farther lower value", arr: []int{5, 6, 1}, d: 2, want: 2},
		{name: "larger distance enables longer path", arr: []int{3, 1, 2}, d: 2, want: 3},
		{name: "distance limit prevents longer path", arr: []int{3, 1, 2}, d: 1, want: 2},
		{name: "choose any starting index", arr: []int{1, 2, 3, 4}, d: 4, want: 4},
		{name: "cannot combine both branches", arr: []int{1, 3, 2}, d: 1, want: 2},
		{name: "maximum height", arr: []int{100000, 1}, d: 2, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := append([]int(nil), tt.arr...)
			if got := maxJumps(arr, tt.d); got != tt.want {
				t.Errorf("maxJumps(%v, %d) = %d, want %d", tt.arr, tt.d, got, tt.want)
			}
		})
	}
}

func TestMaxJumpsLargeInput(t *testing.T) {
	tests := []struct {
		name  string
		equal bool
		want  int
	}{
		{name: "descending chain", equal: false, want: 1000},
		{name: "all equal", equal: true, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, 1000)
			for i := range arr {
				if tt.equal {
					arr[i] = 100000
				} else {
					arr[i] = len(arr) - i
				}
			}
			if got := maxJumps(arr, 1000); got != tt.want {
				t.Errorf("maxJumps(1000 elements, %s) = %d, want %d", tt.name, got, tt.want)
			}
		})
	}
}
