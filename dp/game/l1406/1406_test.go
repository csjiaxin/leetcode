package l1406

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestStoneGameIII(t *testing.T) {
	for _, tt := range []struct {
		a    []int
		want string
	}{
		{nil, "Tie"}, {[]int{1, 2, 3, 7}, "Bob"}, {[]int{1, 2, 3, -9}, "Alice"},
		{[]int{1, 2, 3, 6}, "Tie"}, {[]int{-1, -2, -3}, "Tie"},
		{[]int{-1}, "Bob"}, {[]int{5}, "Alice"}, {[]int{0, 0, 0, 0}, "Tie"},
	} {
		before := append([]int(nil), tt.a...)
		for call := 0; call < 2; call++ {
			if got := stoneGameIII(tt.a); got != tt.want {
				t.Fatalf("%v got %s want %s", tt.a, got, tt.want)
			}
		}
		if !reflect.DeepEqual(before, tt.a) {
			t.Fatal("input mutated")
		}
	}
}

// 显式先后手极大极小，得分始终以 Alice 减 Bob 表示。
func bruteIII(a []int, start int, alice bool) int {
	if start == len(a) {
		return 0
	}
	best, initialized, sum := 0, false, 0
	for x := 1; x <= 3 && start+x <= len(a); x++ {
		sum += a[start+x-1]
		candidate := bruteIII(a, start+x, !alice)
		if alice {
			candidate += sum
		} else {
			candidate -= sum
		}
		if !initialized || (alice && candidate > best) || (!alice && candidate < best) {
			best, initialized = candidate, true
		}
	}
	return best
}
func TestStoneIIIAgainstMinimax(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	for trial := 0; trial < 200; trial++ {
		a := make([]int, 1+rng.Intn(8))
		for i := range a {
			a[i] = rng.Intn(11) - 5
		}
		diff := bruteIII(a, 0, true)
		want := "Tie"
		if diff > 0 {
			want = "Alice"
		} else if diff < 0 {
			want = "Bob"
		}
		if got := stoneGameIII(a); got != want {
			t.Fatalf("%v got %s want %s", a, got, want)
		}
	}
}
func TestStoneIIILarge(t *testing.T) {
	a := make([]int, 50000)
	if got := stoneGameIII(a); got != "Tie" {
		t.Fatal(got)
	}
}
