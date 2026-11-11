package pqueue

import (
	"math/rand"
	"slices"
	"testing"
)

func TestQueue(t *testing.T) {
	for _, desc := range []bool{false, true} {
		less := func(a, b int) bool {
			if desc {
				return a > b
			}
			return a < b
		}
		r := rand.New(rand.NewSource(1))
		input := make([]int, 1000)
		for i := range input {
			input[i] = r.Intn(100)
		}
		q := New(less, input[:500]...)
		for _, x := range input[500:] {
			q.Push(x)
		}
		want := slices.Clone(input)
		slices.Sort(want)
		if desc {
			slices.Reverse(want)
		}
		for i, x := range want {
			if q.Len() != len(want)-i || q.Peek() != x || q.Pop() != x {
				t.Fatalf("%s: order/length mismatch desc=%t at %d", t.Name(), desc, i)
			}
		}
		if q.Len() != 0 {
			t.Fatalf("%s: expected empty", t.Name())
		}
	}
}
