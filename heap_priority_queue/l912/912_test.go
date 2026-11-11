package l912

import (
	"math/rand"
	"slices"
	"testing"
)

func TestSortArray(t *testing.T) {
	for _, tt := range []struct {
		name string
		nums []int
	}{
		{"example", []int{5, 2, 3, 1}}, {"duplicates", []int{5, 1, 1, 2, 0, 0}}, {"negative", []int{-5, -1, -10, 0}}, {"single", []int{4}}, {"odd length", []int{4, 3, 2, 1, 5}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			want := slices.Clone(tt.nums)
			slices.Sort(want)
			if got := sortArray(tt.nums); !slices.Equal(got, want) {
				t.Errorf("%s: got %v, want %v", t.Name(), got, want)
			}
		})
	}
}
func TestSortArrayRandom(t *testing.T) {
	r := rand.New(rand.NewSource(912))
	for n := 1; n <= 200; n++ {
		a := make([]int, n)
		for i := range a {
			a[i] = r.Intn(101) - 50
		}
		want := slices.Clone(a)
		slices.Sort(want)
		if got := sortArray(a); !slices.Equal(got, want) {
			t.Fatalf("%s: n=%d got %v want %v", t.Name(), n, got, want)
		}
	}
}
