package main

import "fmt"

var adj [][]int
var visit []bool

func dfs(at int) int64 {
	visit[at] = true
	r := int64(1)
	for _, to := range adj[at] {
		if !visit[to] {
			r += dfs(to)
		}
	}
	return r

}
func countPairs(n int, edges [][]int) int64 {
	adj = make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	visit = make([]bool, n)
	ca := make([]int64, 0)

	for i := 0; i < n; i++ {
		if !visit[i] {
			c := dfs(i)
			ca = append(ca, c)
		}
	}
	var ret int64
	for i := 0; i < len(ca); i++ {
		var s int64
		for j := i + 1; j < len(ca); j++ {
			s += ca[j]
		}
		ret += ca[i] * s
	}
	return ret
}
func main() {
	fmt.Println(countPairs(7, [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}}))
}
