package l1140

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestStoneGameII(t *testing.T) {
	for _, tt := range []struct {
		a    []int
		want int
	}{
		{nil, 0}, {[]int{2, 7, 9, 4, 4}, 10}, {[]int{1, 2, 3, 4, 5, 100}, 104},
		{[]int{8}, 8}, {[]int{1, 2}, 3}, {[]int{1, 1, 1}, 2}, {[]int{1, 1, 1, 1}, 2},
	} {
		before := append([]int(nil), tt.a...)
		for call := 0; call < 2; call++ {
			if got := stoneGameII(tt.a); got != tt.want {
				t.Fatalf("%v got %d want %d", tt.a, got, tt.want)
			}
		}
		if !reflect.DeepEqual(before, tt.a) {
			t.Fatal("input mutated")
		}
	}
}

// 直接最大化/最小化 Alice 的得分，独立验证分差还原公式。
func bruteII(a []int, start, m int, alice bool) int {
	if start == len(a) {
		return 0
	}
	best, initialized, sum := 0, false, 0
	for x := 1; x <= 2*m && start+x <= len(a); x++ {
		sum += a[start+x-1]
		nextM := m
		if x > nextM {
			nextM = x
		}
		candidate := bruteII(a, start+x, nextM, !alice)
		if alice {
			candidate += sum
		}
		if !initialized || (alice && candidate > best) || (!alice && candidate < best) {
			best, initialized = candidate, true
		}
	}
	return best
}
func TestStoneIIAgainstMinimax(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	for trial := 0; trial < 200; trial++ {
		a := make([]int, 1+rng.Intn(10))
		for i := range a {
			a[i] = 1 + rng.Intn(10)
		}
		want := bruteII(a, 0, 1, true)
		if got := stoneGameII(a); got != want {
			t.Fatalf("%v got %d want %d", a, got, want)
		}
	}
}
