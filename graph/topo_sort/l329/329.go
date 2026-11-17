package l329

// Kahn 不修改 indegree 和 seeds，返回处理的节点数。
// seeds 必须互不重复且入度为零，next 必须同步枚举每条出边一次。
// indegree 必须与这些边的数量一致（包括重边），节点编号必须合法。
// relax 在递减后继入度前调用；onPop 在全部前置贡献就绪后调用。
// DP 初值和边上的转移由调用方提供，回调不能修改图结构。
func Kahn(indegree []int, seeds []int, next func(u int, yield func(v int)), onPop func(u int), relax func(u, v int)) int {
	remaining := append([]int(nil), indegree...)
	queue := append([]int(nil), seeds...)
	for head := 0; head < len(queue); head++ {
		u := queue[head]
		if onPop != nil {
			onPop(u)
		}
		next(u, func(v int) {
			if relax != nil {
				relax(u, v)
			}
			remaining[v]--
			if remaining[v] == 0 {
				queue = append(queue, v)
			}
		})
	}
	return len(queue)
}

// ZeroIndegree 收集普通 DAG 遍历的全部源点；不修改入度。
// 依赖外部 supplies 的题目应由调用方显式构造 seeds。
func ZeroIndegree(indegree []int) []int {
	seeds := make([]int, 0)
	for u, degree := range indegree {
		if degree == 0 {
			seeds = append(seeds, u)
		}
	}
	return seeds
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
