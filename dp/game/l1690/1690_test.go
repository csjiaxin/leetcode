package l1690

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestStoneGameVII(t *testing.T) {
	for _, tt := range []struct {
		a    []int
		want int
	}{
		{nil, 0}, {[]int{5}, 0}, {[]int{5, 3, 1, 4, 2}, 6},
		{[]int{7, 90, 5, 1, 100, 10, 10, 2}, 122}, {[]int{1, 2}, 2}, {[]int{1, 1, 1}, 1},
	} {
		before := append([]int(nil), tt.a...)
		for call := 0; call < 2; call++ {
			if got := stoneGameVII(tt.a); got != tt.want {
				t.Fatalf("%v got %d want %d", tt.a, got, tt.want)
			}
		}
		if !reflect.DeepEqual(before, tt.a) {
			t.Fatal("input mutated")
		}
	}
}
func bruteVII(a []int, alice bool) int {
	if len(a) <= 1 {
		return 0
	}
	sum := 0
	for _, v := range a {
		sum += v
	}
	left, right := bruteVII(a[1:], !alice), bruteVII(a[:len(a)-1], !alice)
	if alice {
		left += sum - a[0]
		right += sum - a[len(a)-1]
		if left > right {
			return left
		}
		return right
	}
	left -= sum - a[0]
	right -= sum - a[len(a)-1]
	if left < right {
		return left
	}
	return right
}
func TestStoneVIIAgainstMinimax(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	for trial := 0; trial < 200; trial++ {
		a := make([]int, 1+rng.Intn(9))
		for i := range a {
			a[i] = 1 + rng.Intn(20)
		}
		want := bruteVII(a, true)
		if got := stoneGameVII(a); got != want {
			t.Fatalf("%v got %d want %d", a, got, want)
		}
	}
}
