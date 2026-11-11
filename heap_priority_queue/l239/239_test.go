package l239

import (
	"math/rand"
	"slices"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"example", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}}, {"single", []int{1}, 1, []int{1}}, {"duplicates", []int{2, 2, 2}, 2, []int{2, 2}}, {"descending", []int{5, 4, 3, 2}, 2, []int{5, 4, 3}}, {"increasing", []int{1, 2, 3, 4}, 2, []int{2, 3, 4}}, {"negative full", []int{-3, -1, -2}, 3, []int{-1}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.nums, tt.k); !slices.Equal(got, tt.want) {
				t.Errorf("%s: got %v want %v", t.Name(), got, tt.want)
			}
		})
	}
}
func TestMaxSlidingWindowRandom(t *testing.T) {
	r := rand.New(rand.NewSource(239))
	for trial := 0; trial < 300; trial++ {
		n := r.Intn(60) + 1
		k := r.Intn(n) + 1
		nums := make([]int, n)
		for i := range nums {
			nums[i] = r.Intn(21) - 10
		}
		want := make([]int, n-k+1)
		for i := range want {
			want[i] = slices.Max(nums[i : i+k])
		}
		if got := maxSlidingWindow(nums, k); !slices.Equal(got, want) {
			t.Fatalf("%s: trial %d got %v want %v", t.Name(), trial, got, want)
		}
	}
}
