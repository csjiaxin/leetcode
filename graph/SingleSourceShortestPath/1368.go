package main

import (
	"fmt"
	"math"
)

type E struct {
	i int
	j int
}

func dfs(cost, i, j, m, n int, grid [][]int, q *[]E, dp [][]int) {
	if i < 0 || i >= m || j < 0 || j >= n || dp[i][j] != math.MaxInt {
		return
	}
	*q = append(*q, E{i, j})
	dp[i][j] = cost
	switch grid[i][j] {
	case 1:
		dfs(cost, i, j+1, m, n, grid, q, dp)
	case 2:
		dfs(cost, i, j-1, m, n, grid, q, dp)
	case 3:
		dfs(cost, i+1, j, m, n, grid, q, dp)
	case 4:
		dfs(cost, i-1, j, m, n, grid, q, dp)
	}
}

var dir4 = [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

func minCost(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}
	cost := 0
	q := []E{{0, 0}}
	dfs(cost, 0, 0, m, n, grid, &q, dp)
	for len(q) > 0 {
		cost += 1
		size := len(q)
		for i := 0; i < size; i++ {
			front := q[0]
			q = q[1:]
			for j := 0; j < 4; j++ {
				dfs(cost, front.i+dir4[j][0], front.j+dir4[j][1], m, n, grid, &q, dp)
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {
	//3
	fmt.Println(minCost([][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}))
	// fmt.Println(minCost([][]int{{1, 1, 3}, {3, 2, 2}, {1, 1, 4}}))
}
