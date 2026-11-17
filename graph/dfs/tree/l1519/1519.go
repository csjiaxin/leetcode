package main

import "fmt"

var adj [][]int
var visit []bool
var result []int

func countSubTrees(n int, edges [][]int, labels string) []int {
	adj = make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	result = make([]int, n)
	visit = make([]bool, n)
	dfs(0, labels)
	return result
}
func dfs(at int, labels string) map[byte]int {
	visit[at] = true
	ret := map[byte]int{}
	for _, to := range adj[at] {
		if !visit[to] {
			toret := dfs(to, labels)
			for k, v := range toret {
				ret[k] += v
			}

		}
	}
	ret[labels[at]] += 1
	result[at] = ret[labels[at]]

	return ret
}
func main() {
	fmt.Println(countSubTrees(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, "abaedcd"))
}
