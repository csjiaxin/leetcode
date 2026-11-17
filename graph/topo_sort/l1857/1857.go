package l1857

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
