package l33

import "testing"

func TestSearch(t *testing.T) {
	for _, tt := range []struct {
		name         string
		nums         []int
		target, want int
	}{
		{"example", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4}, {"missing", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"empty", nil, 1, -1}, {"single", []int{1}, 1, 0}, {"two", []int{3, 1}, 1, 1},
		{"sorted", []int{1, 2, 3}, 3, 2}, {"left", []int{4, 5, 6, 1, 2, 3}, 5, 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestAllRotations(t *testing.T) {
	for n := 1; n <= 20; n++ {
		for pivot := 0; pivot < n; pivot++ {
			nums := make([]int, n)
			for i := range nums {
				nums[i] = (pivot + i) % n
			}
			for target := -1; target <= n; target++ {
				want := -1
				for i, v := range nums {
					if v == target {
						want = i
					}
				}
				if got := search(nums, target); got != want {
					t.Fatalf("nums=%v target=%d got=%d want=%d", nums, target, got, want)
				}
			}
		}
	}
}
