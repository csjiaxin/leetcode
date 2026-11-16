package l1462

// Kahn 按依赖顺序处理节点；seeds 不重复且入度为零，next 同步枚举后继。
// 返回处理数；入度和 seeds 均不修改，DP 状态由调用方初始化。
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

func ZeroIndegree(indegree []int) []int {
	seeds := make([]int, 0)
	for u, degree := range indegree {
		if degree == 0 {
			seeds = append(seeds, u)
		}
	}
	return seeds
}

// https://leetcode.com/problems/course-schedule-iv/
// 题设保证依赖图无环，before[v][x] 表示 x 是 v 的祖先。
func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	before := make([][]bool, numCourses)
	for v := range before {
		before[v] = make([]bool, numCourses)
	}
	for _, e := range prerequisites {
		graph[e[0]] = append(graph[e[0]], e[1])
		indegree[e[1]]++
	}
	Kahn(indegree, ZeroIndegree(indegree),
		func(u int, yield func(int)) {
			for _, v := range graph[u] {
				yield(v)
			}
		}, nil,
		func(u, v int) {
			before[v][u] = true
			for x := 0; x < numCourses; x++ {
				before[v][x] = before[v][x] || before[u][x]
			}
		},
	)
	answer := make([]bool, len(queries))
	for i, q := range queries {
		answer[i] = before[q[1]][q[0]]
	}
	return answer
}
