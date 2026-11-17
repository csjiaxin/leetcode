package main

import "fmt"

// https://leetcode.com/problems/min-cost-to-connect-all-points/
func minCostConnectPoints(points [][]int) int {
	return DensePrim(len(points), func(i, j int) int {
		dx := points[i][0] - points[j][0]
		if dx < 0 {
			dx = -dx
		}
		dy := points[i][1] - points[j][1]
		if dy < 0 {
			dy = -dy
		}
		return dx + dy
	})
}

// DensePrim returns the MST cost of a complete undirected graph.
// cost(i, j) must be symmetric, and all edge costs and their sum must fit in int.
func DensePrim(n int, cost func(int, int) int) int {
	if n <= 1 {
		return 0
	}
	included := make([]bool, n)
	key := make([]int, n)
	included[0] = true
	for vertex := 1; vertex < n; vertex++ {
		key[vertex] = cost(0, vertex)
	}
	total := 0
	for count := 1; count < n; count++ {
		next := -1
		for vertex := 0; vertex < n; vertex++ {
			if !included[vertex] && (next == -1 || key[vertex] < key[next]) {
				next = vertex
			}
		}
		included[next] = true
		total += key[next]
		// key[v] is the cheapest single edge from the current tree to v.
		for vertex := 0; vertex < n; vertex++ {
			if !included[vertex] {
				weight := cost(next, vertex)
				if weight < key[vertex] {
					key[vertex] = weight
				}
			}
		}
	}
	return total
}

func main() {
	fmt.Println(minCostConnectPoints([][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}}))
}
