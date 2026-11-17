package l329

import topotemplate "leetcode/graph/topo_sort/template"

// Kahn 按依赖顺序处理节点；seeds 不重复且入度为零，next 同步枚举后继。
// 返回处理数；入度和 seeds 均不修改，DP 状态由调用方初始化。
func Kahn(indegree []int, seeds []int, next func(u int, yield func(v int)), onPop func(u int), relax func(u, v int)) int {
	return topotemplate.Kahn(indegree, seeds, next, onPop, relax)
}

func ZeroIndegree(indegree []int) []int {
	return topotemplate.ZeroIndegree(indegree)
}

// https://leetcode.com/problems/longest-increasing-path-in-a-matrix/
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	n := rows * cols
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	// 严格递增边隐式构成 DAG，不修改矩阵，不存邻接表。
	next := func(u int, yield func(int)) {
		r, c := u/cols, u%cols
		for _, d := range directions {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && matrix[nr][nc] > matrix[r][c] {
				yield(nr*cols + nc)
			}
		}
	}
	indegree := make([]int, n)
	for u := 0; u < n; u++ {
		next(u, func(v int) { indegree[v]++ })
	}
	length := make([]int, n)
	for u := range length {
		length[u] = 1
	}
	answer := 0
	Kahn(indegree, ZeroIndegree(indegree), next,
		func(u int) {
			if length[u] > answer {
				answer = length[u]
			}
		},
		func(u, v int) {
			if length[u]+1 > length[v] {
				length[v] = length[u] + 1
			}
		},
	)
	return answer
}
