package segment

import (
	"math/rand"
	"testing"
)

func TestNumArrayExample(t *testing.T) {
	obj := Constructor([]int{1, 3, 5})
	if got := obj.SumRange(0, 2); got != 9 {
		t.Fatalf("SumRange(0,2) = %d, want 9", got)
	}
	obj.Update(1, 2)
	if got := obj.SumRange(0, 2); got != 8 {
		t.Fatalf("SumRange(0,2) after update = %d, want 8", got)
	}
}

func TestNumArrayCases(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		left, right int
		want        int
	}{
		{"single element", []int{7}, 0, 0, 7},
		{"single index in middle", []int{1, 2, 3, 4, 5}, 2, 2, 3},
		{"full range non power of two", []int{1, 2, 3, 4, 5}, 0, 4, 15},
		{"prefix", []int{1, 2, 3, 4, 5}, 0, 2, 6},
		{"suffix", []int{1, 2, 3, 4, 5}, 3, 4, 9},
		{"negatives", []int{-2, 0, 3, -5, 2, -1}, 0, 5, -3},
		{"negatives sub range", []int{-2, 0, 3, -5, 2, -1}, 2, 5, -1},
		{"power of two length", []int{1, 1, 1, 1, 1, 1, 1, 1}, 1, 6, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := Constructor(tt.nums)
			if got := obj.SumRange(tt.left, tt.right); got != tt.want {
				t.Errorf("SumRange(%d,%d) = %d, want %d", tt.left, tt.right, got, tt.want)
			}
		})
	}
}

func TestNumArrayUpdateEdges(t *testing.T) {
	obj := Constructor([]int{1, 2, 3})
	obj.Update(0, 10)
	obj.Update(2, -3)
	if got := obj.SumRange(0, 0); got != 10 {
		t.Errorf("SumRange(0,0) = %d, want 10", got)
	}
	if got := obj.SumRange(2, 2); got != -3 {
		t.Errorf("SumRange(2,2) = %d, want -3", got)
	}
	if got := obj.SumRange(0, 2); got != 9 {
		t.Errorf("SumRange(0,2) = %d, want 9", got)
	}
}

// TestNumArrayRandom compares against a brute-force implementation.
func TestNumArrayRandom(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	for iter := 0; iter < 200; iter++ {
		n := r.Intn(50) + 1
		nums := make([]int, n)
		for i := range nums {
			nums[i] = r.Intn(201) - 100
		}
		obj := Constructor(nums)
		ref := append([]int(nil), nums...)
		for op := 0; op < 100; op++ {
			if r.Intn(2) == 0 {
				idx, val := r.Intn(n), r.Intn(201)-100
				obj.Update(idx, val)
				ref[idx] = val
				continue
			}
			l := r.Intn(n)
			rr := l + r.Intn(n-l)
			want := 0
			for i := l; i <= rr; i++ {
				want += ref[i]
			}
			if got := obj.SumRange(l, rr); got != want {
				t.Fatalf("n=%d SumRange(%d,%d) = %d, want %d", n, l, rr, got, want)
			}
		}
	}
}
