package l1857

import topotemplate "leetcode/graph/topo_sort/template"

// Kahn 按依赖顺序处理节点；seeds 不重复且入度为零，next 同步枚举后继。
// 返回处理数；入度和 seeds 均不修改，DP 状态由调用方初始化。
func Kahn(indegree []int, seeds []int, next func(u int, yield func(v int)), onPop func(u int), relax func(u, v int)) int {
	return topotemplate.Kahn(indegree, seeds, next, onPop, relax)
}

func ZeroIndegree(indegree []int) []int {
	return topotemplate.ZeroIndegree(indegree)
}

// https://leetcode.com/problems/largest-color-value-in-a-directed-graph/
func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	graph := make([][]int, n)
	indegree := make([]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		indegree[e[1]]++
	}
	// dp[v][c]：以 v 为终点的路径中，颜色 c 出现次数的最大值。
	dp := make([][26]int, n)
	for v := range dp {
		dp[v][colors[v]-'a'] = 1
	}
	answer := 0
	processed := Kahn(indegree, ZeroIndegree(indegree),
		func(u int, yield func(int)) {
			for _, v := range graph[u] {
				yield(v)
			}
		},
		func(u int) {
			for _, value := range dp[u] {
				if value > answer {
					answer = value
				}
			}
		},
		func(u, v int) {
			for c := 0; c < 26; c++ {
				candidate := dp[u][c]
				if c == int(colors[v]-'a') {
					candidate++
				}
				if candidate > dp[v][c] {
					dp[v][c] = candidate
				}
			}
		},
	)
	if processed != n {
		return -1
	}
	return answer
}
