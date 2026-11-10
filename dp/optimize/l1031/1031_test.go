package l1031

import "testing"

func TestMaxSumTwoNoOverlap(t *testing.T) {
	tests := []struct {
		name                string
		nums                []int
		firstLen, secondLen int
		want                int
	}{
		{name: "example one", nums: []int{0, 6, 5, 2, 2, 5, 1, 9, 4}, firstLen: 1, secondLen: 2, want: 20},
		{name: "example two", nums: []int{3, 8, 1, 3, 2, 1, 8, 9, 0}, firstLen: 3, secondLen: 2, want: 29},
		{name: "example three", nums: []int{2, 1, 5, 6, 0, 9, 5, 0, 3, 8}, firstLen: 4, secondLen: 3, want: 31},
		{name: "minimum input", nums: []int{1, 2}, firstLen: 1, secondLen: 1, want: 3},
		{name: "all zero", nums: []int{0, 0, 0, 0}, firstLen: 1, secondLen: 2, want: 0},
		{name: "equal lengths", nums: []int{5, 5, 0, 5, 5}, firstLen: 2, secondLen: 2, want: 20},
		{name: "windows cover entire input", nums: []int{1, 2, 3, 4, 5}, firstLen: 2, secondLen: 3, want: 15},
		{name: "first length on left", nums: []int{10, 10, 0, 0, 9}, firstLen: 2, secondLen: 1, want: 29},
		{name: "second length on left", nums: []int{9, 0, 0, 10, 10}, firstLen: 2, secondLen: 1, want: 29},
		{name: "do not overlap largest windows", nums: []int{0, 10, 10, 10, 0}, firstLen: 2, secondLen: 2, want: 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumTwoNoOverlap(tt.nums, tt.firstLen, tt.secondLen); got != tt.want {
				t.Errorf("maxSumTwoNoOverlap(%v, %d, %d) = %d, want %d", tt.nums, tt.firstLen, tt.secondLen, got, tt.want)
			}
		})
	}
}

func TestMaxSumTwoNoOverlapAgainstBruteForce(t *testing.T) {
	// Exhaust all arrays of length 2 through 6 with values in [0, 2].
	for n := 2; n <= 6; n++ {
		count := 1
		for i := 0; i < n; i++ {
			count *= 3
		}
		for code := 0; code < count; code++ {
			nums := make([]int, n)
			x := code
			for i := range nums {
				nums[i] = x % 3
				x /= 3
			}
			for firstLen := 1; firstLen < n; firstLen++ {
				for secondLen := 1; firstLen+secondLen <= n; secondLen++ {
					want := 0
					for i := 0; i+firstLen <= n; i++ {
						for j := 0; j+secondLen <= n; j++ {
							if i+firstLen > j && j+secondLen > i {
								continue
							}
							sum := 0
							for _, v := range nums[i : i+firstLen] {
								sum += v
							}
							for _, v := range nums[j : j+secondLen] {
								sum += v
							}
							want = max(want, sum)
						}
					}
					if got := maxSumTwoNoOverlap(nums, firstLen, secondLen); got != want {
						t.Fatalf("maxSumTwoNoOverlap(%v, %d, %d) = %d, want %d", nums, firstLen, secondLen, got, want)
					}
				}
			}
		}
	}
}
