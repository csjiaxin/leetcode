package main

import "fmt"

type Edge struct {
	From, To int
	Weight   int64
}

// FloydWarshall 返回有向图全源最短距离及可达性；边端点须在 [0,n) 内。
func FloydWarshall(n int, edges []Edge) ([][]int64, [][]bool, error) {
	if n < 0 {
		return nil, nil, fmt.Errorf("n must be nonnegative")
	}
	dist := make([][]int64, n)
	reachable := make([][]bool, n)
	for i := range dist {
		dist[i] = make([]int64, n)
		reachable[i] = make([]bool, n)
		reachable[i][i] = true
	}
	for _, edge := range edges {
		u, v := edge.From, edge.To
		if u < 0 || u >= n || v < 0 || v >= n {
			return nil, nil, fmt.Errorf("edge endpoint out of range")
		}
		if !reachable[u][v] || edge.Weight < dist[u][v] {
			dist[u][v] = edge.Weight
			reachable[u][v] = true
		}
	}
	if !FloydClosure(dist, reachable, checkedAdd, func(a, b int64) bool { return a < b }) {
		return nil, nil, fmt.Errorf("distance overflow")
	}
	for i := range dist {
		if dist[i][i] < 0 {
			return nil, nil, fmt.Errorf("negative cycle detected")
		}
	}
	return dist, reachable, nil
}

// FloydClosure 按允许的中间节点递增；extend 失败时停止并返回 false。
func FloydClosure[T any](value [][]T, reachable [][]bool, extend func(T, T) (T, bool), better func(T, T) bool) bool {
	for k := range value {
		for i := range value {
			if !reachable[i][k] {
				continue
			}
			for j := range value {
				if !reachable[k][j] {
					continue
				}
				candidate, ok := extend(value[i][k], value[k][j])
				if !ok {
					return false
				}
				if !reachable[i][j] || better(candidate, value[i][j]) {
					value[i][j] = candidate
					reachable[i][j] = true
				}
			}
		}
	}
	return true
}

func checkedAdd(a, b int64) (int64, bool) {
	const highest = int64(1<<63 - 1)
	const lowest = int64(-1 << 63)
	if b > 0 && a > highest-b || b < 0 && a < lowest-b {
		return 0, false
	}
	return a + b, true
}
