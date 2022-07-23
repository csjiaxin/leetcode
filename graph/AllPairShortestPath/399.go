package main

import (
	"fmt"
	"math"
)

var adj [][]int

func buildAdjList(N int, connections [][]int) {
	adj := make([][]int, N)
	for i := 0; i < N; i++ {
		adj[i] = make([]int, N)
	}
	for _, e := range connections {
		adj[e[0]][e[1]] = e[2]
		adj[e[1]][e[0]] = e[2]
	}
}
func Init2D(r, c int, v int) [][]int {

	adj := make([][]int, r)
	for i := 0; i < r; i++ {
		adj[i] = make([]int, c)
		for j := 0; j < c; j++ {
			adj[i][j] = v
		}
	}
	return adj
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	n:=
	d := Init2D(n, n, math.MaxInt)
	for i := 0; i < n; i++ {
		d[i][i] = 0
	}
	for _, e := range edges {
		d[e[0]][e[1]] = e[2]
		d[e[1]][e[0]] = e[2]
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if d[i][k] != math.MaxInt && d[k][j] != math.MaxInt {

					d[i][j] = min(d[i][j], d[i][k]+d[k][j])
					d[j][i] = d[i][j]
				}
			}
		}
	}
	minNum := math.MaxInt
	minI := -1
	for i := 0; i < n; i++ {
		num := 0
		for j := 0; j < n; j++ {
			if i != j && d[i][j] <= distanceThreshold {
				num += 1
			}
		}
		if num <= minNum {
			minNum = num
			minI = i
		}
	}
	return minI

}

func main() {
	fmt.Println(findTheCity(5, [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, 2))

}
