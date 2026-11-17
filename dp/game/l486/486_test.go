package l486

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestPredictTheWinner(t *testing.T) {
	for _, tt := range []struct {
		nums []int
		want bool
	}{
		{nil, true}, {[]int{1}, true}, {[]int{1, 5, 2}, false}, {[]int{1, 5, 233, 7}, true},
		{[]int{0, 0}, true}, {[]int{1, 1}, true}, {[]int{1, 2, 1}, true},
	} {
		before := append([]int(nil), tt.nums...)
		for call := 0; call < 2; call++ {
			if got := predictTheWinner(tt.nums); got != tt.want {
				t.Fatalf("%v got %v want %v", tt.nums, got, tt.want)
			}
		}
		if !reflect.DeepEqual(before, tt.nums) {
			t.Fatal("input mutated")
		}
	}
}

// 显式区分先后手：先手最大化自身得分，后手最小化先手得分。
func bruteFirst(a []int, left, right int, first bool) int {
	if left > right {
		return 0
	}
	x, y := 0, 0
	if first {
		x = a[left] + bruteFirst(a, left+1, right, false)
		y = a[right] + bruteFirst(a, left, right-1, false)
		if x > y {
			return x
		}
		return y
	}
	x = bruteFirst(a, left+1, right, true)
	y = bruteFirst(a, left, right-1, true)
	if x < y {
		return x
	}
	return y
}
func TestPredictAgainstMinimax(t *testing.T) {
	rng := rand.New(rand.NewSource(42))
	for trial := 0; trial < 200; trial++ {
		a := make([]int, 1+rng.Intn(9))
		sum := 0
		for i := range a {
			a[i] = rng.Intn(10)
			sum += a[i]
		}
		want := 2*bruteFirst(a, 0, len(a)-1, true) >= sum
		if got := predictTheWinner(a); got != want {
			t.Fatalf("%v got %v want %v", a, got, want)
		}
	}
}
