package l210

import topotemplate "leetcode/graph/topo_sort/template"

// Kahn 按依赖顺序处理节点；seeds 不重复且入度为零，next 同步枚举后继。
// 返回处理数；入度和 seeds 均不修改，DP 状态由调用方初始化。
func Kahn(indegree []int, seeds []int, next func(u int, yield func(v int)), onPop func(u int), relax func(u, v int)) int {
	return topotemplate.Kahn(indegree, seeds, next, onPop, relax)
}

func ZeroIndegree(indegree []int) []int {
	return topotemplate.ZeroIndegree(indegree)
}

// https://leetcode.com/problems/course-schedule-ii/
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
		indegree[p[0]]++
	}
	order := make([]int, 0, numCourses)
	processed := Kahn(indegree, ZeroIndegree(indegree),
		func(u int, yield func(int)) {
			for _, v := range graph[u] {
				yield(v)
			}
		},
		func(u int) { order = append(order, u) }, nil,
	)
	if processed != numCourses {
		return []int{}
	}
	return order
}
