package l480

import (
	"math/rand"
	"slices"
	"testing"
)

func TestMedianSlidingWindow(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
		k    int
		want []float64
	}{
		{"example", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []float64{1, -1, -1, 3, 5, 6}}, {"even", []int{1, 2, 3, 4}, 2, []float64{1.5, 2.5, 3.5}}, {"single windows", []int{3, -1, 2}, 1, []float64{3, -1, 2}}, {"duplicates", []int{2, 2, 2, 2}, 3, []float64{2, 2}}, {"int32 extremes", []int{-2147483648, 2147483647}, 2, []float64{-0.5}}, {"full window", []int{3, 1, 2}, 3, []float64{2}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := medianSlidingWindow(tt.nums, tt.k); !slices.Equal(got, tt.want) {
				t.Errorf("%s: got %v want %v", t.Name(), got, tt.want)
			}
		})
	}
}
func TestMedianSlidingWindowRandom(t *testing.T) {
	r := rand.New(rand.NewSource(480))
	for trial := 0; trial < 500; trial++ {
		n := r.Intn(60) + 1
		k := r.Intn(n) + 1
		nums := make([]int, n)
		for i := range nums {
			nums[i] = r.Intn(21) - 10
		}
		want := make([]float64, n-k+1)
		for i := range want {
			window := slices.Clone(nums[i : i+k])
			slices.Sort(window)
			want[i] = float64(window[k/2])
			if k%2 == 0 {
				want[i] = float64(window[k/2-1])/2 + float64(window[k/2])/2
			}
		}
		if got := medianSlidingWindow(nums, k); !slices.Equal(got, want) {
			t.Fatalf("%s: trial %d nums=%v k=%d got=%v want=%v", t.Name(), trial, nums, k, got, want)
		}
	}
}
