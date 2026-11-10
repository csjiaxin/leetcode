package l1345

import "testing"

func TestMinJumps(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{name: "all distinct", arr: []int{1, 2, 3, 4, 5}, want: 4},
		{name: "example one", arr: []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}, want: 3},
		{name: "example two", arr: []int{7}, want: 0},
		{name: "example three", arr: []int{7, 6, 9, 6, 9, 6, 9, 7}, want: 1},
		{name: "single zero", arr: []int{0}, want: 0},
		{name: "two distinct values", arr: []int{1, 2}, want: 1},
		{name: "two equal values", arr: []int{5, 5}, want: 1},
		{name: "all equal", arr: []int{3, 3, 3, 3, 3}, want: 1},
		{name: "equal value shortcut in middle", arr: []int{1, 2, 3, 2, 4}, want: 3},
		{name: "move left after equal value jump", arr: []int{1, 2, 3, 4, 6, 5, 1, 8, 9, 5}, want: 3},
		{name: "negative and zero values", arr: []int{-1, 0, 2, -1, 3, 0}, want: 2},
		{name: "extreme equal endpoints", arr: []int{-100000000, 100000000, 0, -100000000}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := append([]int(nil), tt.arr...)
			if got := minJumps(arr); got != tt.want {
				t.Errorf("minJumps(%v) = %d, want %d", tt.arr, got, tt.want)
			}
		})
	}
}

func TestMinJumpsLargeInput(t *testing.T) {
	tests := []struct {
		name string
		kind int
		want int
	}{
		{name: "all equal", kind: 0, want: 1},
		{name: "all distinct", kind: 1, want: 49999},
		{name: "large equal group before distinct destination", kind: 2, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arr := make([]int, 50000)
			for i := range arr {
				if tt.kind == 1 {
					arr[i] = i
				} else {
					arr[i] = 7
				}
			}
			if tt.kind == 2 {
				arr[len(arr)-1] = 8
			}
			if got := minJumps(arr); got != tt.want {
				t.Errorf("minJumps(50000 elements, %s) = %d, want %d", tt.name, got, tt.want)
			}
		})
	}
}
