package l973

import "testing"

func TestKClosest(t *testing.T) {
	for _, tt := range []struct {
		name   string
		points [][]int
		k      int
	}{
		{"example", [][]int{{1, 3}, {-2, 2}}, 1}, {"example two", [][]int{{3, 3}, {5, -1}, {-2, 4}}, 2}, {"ties", [][]int{{1, 0}, {0, 1}, {-1, 0}}, 2}, {"origin", [][]int{{0, 0}, {1, 1}}, 1}, {"large coordinates", [][]int{{10000, 10000}, {-10000, 0}, {0, 0}}, 2}, {"duplicates", [][]int{{1, 1}, {1, 1}, {2, 2}}, 2},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := kClosest(tt.points, tt.k)
			if len(got) != tt.k {
				t.Fatalf("%s: length %d", t.Name(), len(got))
			}
			counts := map[[2]int]int{}
			for _, p := range tt.points {
				counts[[2]int{p[0], p[1]}]++
			}
			maxChosen := int64(-1)
			for _, p := range got {
				key := [2]int{p[0], p[1]}
				if counts[key] == 0 {
					t.Fatalf("%s: invalid point %v", t.Name(), p)
				}
				counts[key]--
				d := int64(p[0])*int64(p[0]) + int64(p[1])*int64(p[1])
				maxChosen = max(maxChosen, d)
			}
			for p, n := range counts {
				if n > 0 && int64(p[0])*int64(p[0])+int64(p[1])*int64(p[1]) < maxChosen {
					t.Fatalf("%s: omitted closer point %v", t.Name(), p)
				}
			}
		})
	}
}
