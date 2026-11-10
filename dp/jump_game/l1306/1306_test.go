package l1306

import "testing"

func TestCanReach(t *testing.T) {
	tests := []struct {
		name  string
		arr   []int
		start int
		want  bool
	}{
		{name: "example one", arr: []int{4, 2, 3, 0, 3, 1, 2}, start: 5, want: true},
		{name: "example two", arr: []int{4, 2, 3, 0, 3, 1, 2}, start: 0, want: true},
		{name: "example three", arr: []int{3, 0, 2, 1, 2}, start: 2, want: false},
		{name: "single zero", arr: []int{0}, start: 0, want: true},
		{name: "start already zero", arr: []int{1, 0, 2}, start: 1, want: true},
		{name: "jump right to zero", arr: []int{1, 0}, start: 0, want: true},
		{name: "jump left to zero", arr: []int{0, 1}, start: 1, want: true},
		{name: "cycle without zero", arr: []int{1, 1}, start: 0, want: false},
		{name: "unreachable zero behind cycle", arr: []int{2, 1, 2, 0}, start: 0, want: false},
		{name: "both jumps out of bounds", arr: []int{0, 2, 1}, start: 1, want: false},
		{name: "exact jumps not maximum lengths", arr: []int{2, 0, 2}, start: 0, want: false},
		{name: "left branch succeeds", arr: []int{0, 1, 2, 3, 4}, start: 2, want: true},
		{name: "reverse direction to reach zero", arr: []int{2, 0, 1, 2, 1}, start: 0, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Give each call its own array to allow in-place visited marking.
			arr := append([]int(nil), tt.arr...)
			if got := canReach(arr, tt.start); got != tt.want {
				t.Errorf("canReach(%v, %d) = %t, want %t", tt.arr, tt.start, got, tt.want)
			}
		})
	}
}

func TestCanReachLargeInput(t *testing.T) {
	for _, reachable := range []bool{true, false} {
		name := "reachable zero at end"
		if !reachable {
			name = "no zero"
		}
		t.Run(name, func(t *testing.T) {
			arr := make([]int, 50000)
			for i := range arr {
				arr[i] = 1
			}
			if reachable {
				arr[len(arr)-1] = 0
			}
			if got := canReach(arr, 0); got != reachable {
				t.Errorf("canReach(50000 elements, %s) = %t, want %t", name, got, reachable)
			}
		})
	}
}
