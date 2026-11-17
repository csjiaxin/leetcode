package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = math.MaxInt
	}
	dist[src] = 0
	for i := 0; i <= k; i++ {
		lastdist := make([]int, n)
		copy(lastdist, dist)
		for _, e := range flights {
			if lastdist[e[0]] != math.MaxInt && lastdist[e[0]]+e[2] < dist[e[1]] {
				dist[e[1]] = lastdist[e[0]] + e[2]
			}
		}

	}
	m := dist[dst]
	if m == math.MaxInt {
		return -1
	}
	return m

}

func main() {
	fmt.Println(findCheapestPrice(4, [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}, 0, 3, 1))
}
