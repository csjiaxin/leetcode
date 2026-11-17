package binarysearch

import "testing"

func TestFirstTrueEveryBoundary(t *testing.T) {
	for n := 0; n <= 32; n++ {
		for boundary := 0; boundary <= n; boundary++ {
			got := FirstTrue(n, func(i int) bool {
				if i < 0 || i >= n {
					t.Fatalf("predicate index %d outside [0,%d)", i, n)
				}
				return i >= boundary
			})
			if got != boundary {
				t.Fatalf("n=%d boundary=%d: got %d", n, boundary, got)
			}
		}
	}
}
