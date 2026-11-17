package main

import (
	"fmt"
	"math"
)

type Edge struct {
	From, To int
	Weight   int64
}

// checkedAdd prevents both positive overflow and negative underflow.
func checkedAdd(a, b int64) (int64, bool) {
	if b > 0 && a > math.MaxInt64-b {
		return 0, false
	}
	if b < 0 && a < math.MinInt64-b {
		return 0, false
	}
	return a + b, true
}

// Vertices must be numbered 0 through n-1.
// reachable[v] indicates whether dist[v] is meaningful.
func BellmanFord(
	n int,
	edges []Edge,
	src int,
) (dist []int64, reachable []bool, err error) {
	dist = make([]int64, n)
	reachable = make([]bool, n)
	reachable[src] = true

	for pass := 1; pass <= n; pass++ {
		changed := false

		for _, edge := range edges {
			if !reachable[edge.From] {
				continue
			}

			candidate, ok := checkedAdd(dist[edge.From], edge.Weight)
			if !ok {
				return nil, nil, fmt.Errorf("distance overflow")
			}

			if !reachable[edge.To] || candidate < dist[edge.To] {
				if pass == n {
					return nil, nil, fmt.Errorf(
						"negative cycle reachable from source",
					)
				}

				dist[edge.To] = candidate
				reachable[edge.To] = true
				changed = true
			}
		}

		// No improvements means shortest distances are settled.
		if !changed {
			break
		}
	}

	return dist, reachable, nil
}

func main() {
	edges := []Edge{
		{From: 0, To: 1, Weight: 4},
		{From: 0, To: 2, Weight: 5},
		{From: 1, To: 2, Weight: -2},
		{From: 2, To: 3, Weight: 3},
	}

	dist, reachable, err := BellmanFord(4, edges, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(dist)      // [0 4 2 5]
	fmt.Println(reachable) // [true true true true]
}
