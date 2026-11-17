package l2050

import topotemplate "leetcode/graph/topo_sort/template"

// Kahn 按依赖顺序处理节点；seeds 不重复且入度为零，next 同步枚举后继。
// 返回处理数；入度和 seeds 均不修改，DP 状态由调用方初始化。
func Kahn(indegree []int, seeds []int, next func(u int, yield func(v int)), onPop func(u int), relax func(u, v int)) int {
	return topotemplate.Kahn(indegree, seeds, next, onPop, relax)
}

func ZeroIndegree(indegree []int) []int {
	return topotemplate.ZeroIndegree(indegree)
}

// https://leetcode.com/problems/parallel-courses-iii/
func minimumTime(n int, relations [][]int, time []int) int {
	graph := make([][]int, n)
	indegree := make([]int, n)
	for _, e := range relations {
		u, v := e[0]-1, e[1]-1
		graph[u] = append(graph[u], v)
		indegree[v]++
	}
	finish := make([]int64, n)
	for v, duration := range time {
		finish[v] = int64(duration)
	}
	var answer int64
	processed := Kahn(indegree, ZeroIndegree(indegree),
		func(u int, yield func(int)) {
			for _, v := range graph[u] {
				yield(v)
			}
		},
		func(u int) {
			if finish[u] > answer {
				answer = finish[u]
			}
		},
		func(u, v int) {
			// 所有前置完成后才能开始，因此取最大完成时间。
			candidate := finish[u] + int64(time[v])
			if candidate > finish[v] {
				finish[v] = candidate
			}
		},
	)
	if processed != n {
		return -1
	}
	return int(answer)
}
