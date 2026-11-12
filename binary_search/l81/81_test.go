package l81

import "testing"

func TestSearch(t *testing.T) {
	for _, tt := range []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{"example", []int{2, 5, 6, 0, 0, 1, 2}, 0, true}, {"missing", []int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{"ambiguous right", []int{1, 0, 1, 1, 1}, 0, true}, {"ambiguous left", []int{1, 1, 1, 0, 1}, 0, true},
		{"equal present", []int{1, 1, 1}, 1, true}, {"equal absent", []int{1, 1, 1}, 0, false},
		{"empty", nil, 1, false}, {"single", []int{1}, 1, true}, {"two", []int{3, 1}, 1, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}

func TestExhaustiveRotations(t *testing.T) {
	// Enumerate nondecreasing arrays over {0,1,2}, then every rotation.
	for n := 1; n <= 7; n++ {
		a := make([]int, n)
		var visit func(int, int)
		visit = func(i, minimum int) {
			if i < n {
				for v := minimum; v < 3; v++ {
					a[i] = v
					visit(i+1, v)
				}
				return
			}
			for pivot := 0; pivot < n; pivot++ {
				nums := append(append([]int{}, a[pivot:]...), a[:pivot]...)
				for target := -1; target <= 3; target++ {
					want := false
					for _, v := range a {
						if v == target {
							want = true
						}
					}
					if got := search(nums, target); got != want {
						t.Fatalf("nums=%v target=%d got=%v want=%v", nums, target, got, want)
					}
				}
			}
		}
		visit(0, 0)
	}
}
