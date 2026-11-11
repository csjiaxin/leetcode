package l857

import (
	"math"
	"math/bits"
	"math/rand"
	"testing"
)

func TestHiringAgainstSubsets(t *testing.T) {
	r := rand.New(rand.NewSource(857))
	for trial := 0; trial < 100; trial++ {
		n := r.Intn(7) + 1
		k := r.Intn(n) + 1
		quality, wage := make([]int, n), make([]int, n)
		for i := range quality {
			quality[i] = r.Intn(20) + 1
			wage[i] = r.Intn(100) + 1
		}
		want := math.Inf(1)
		for mask := 1; mask < 1<<n; mask++ {
			if bits.OnesCount(uint(mask)) != k {
				continue
			}
			sum := 0
			ratio := 0.0
			for i := 0; i < n; i++ {
				if mask>>i&1 != 0 {
					sum += quality[i]
					ratio = math.Max(ratio, float64(wage[i])/float64(quality[i]))
				}
			}
			want = math.Min(want, float64(sum)*ratio)
		}
		if got := mincostToHireWorkers(quality, wage, k); math.Abs(got-want) > 1e-7 {
			t.Fatalf("%s: trial %d got %v want %v", t.Name(), trial, got, want)
		}
	}
}
