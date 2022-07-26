package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

var presum []int
var demo [][]int

func stoneGameII(piles []int) int {
	n := len(piles)
	presum = make([]int, n+1)
	demo = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		demo[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			demo[i][j] = -1
		}
	}
	presum[n-1] = piles[n-1]
	for i := n - 2; i >= 0; i-- {
		presum[i] = piles[i] + presum[i+1]
	}
	return dfs(piles, 1, 0)
}
func dfs(p []int, m, at int) int {
	if at+2*m >= len(p) {
		return presum[at]
	}
	if demo[at][m] >= 0 {
		return demo[at][m]
	}
	ret := 0
	for x := 1; x <= 2*m; x++ {
		ret = max(ret, presum[at]-dfs(p, max(m, x), at+x))
	}
	demo[at][m] = ret
	return ret
}

func main() {
	fmt.Println(stoneGameII([]int{2, 7, 9, 4, 4}))
}
