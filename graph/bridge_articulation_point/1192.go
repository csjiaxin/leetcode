package main

import "fmt"

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func insert_edge(adj map[int]map[int]bool, from, to int) {
	if v, ok := adj[from]; ok {
		v[to] = true
	} else {
		adj[from] = make(map[int]bool)
		adj[from][to] = true
	}

}
func buildAdj(edges [][]int) map[int]map[int]bool {
	adj := make(map[int]map[int]bool)
	for _, e := range edges {
		insert_edge(adj, e[0], e[1])
		insert_edge(adj, e[1], e[0])
	}
	return adj
}

var bridge [][]int
var ids []int
var lowLink []int
var adj map[int]map[int]bool
var ID int

func dfs(at int, from int, visit []bool) {
	if visit[at] {
		return
	}
	visit[at] = true
	ids[at] = ID
	ID += 1
	lowLink[at] = ids[at]

	for to, _ := range adj[at] {
		if to == from {
			continue
		}
		dfs(to, at, visit)
		lowLink[at] = min(lowLink[at], lowLink[to])
		if ids[at] < lowLink[to] {
			bridge = append(bridge, []int{at, to})
		}

	}

}

func criticalConnections(N int, connections [][]int) [][]int {
	ID = 0
	bridge = [][]int{}
	ids = make([]int, N)
	lowLink = make([]int, N)

	adj = buildAdj(connections)
	visit := make([]bool, N)

	for i := 0; i < N; i++ {
		dfs(i, -1, visit)
	}
	return bridge
}

func main() {
	fmt.Println(criticalConnections(4, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}))
}
