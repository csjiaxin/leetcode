package main

import "fmt"

func FloydWarshall(
	n int,
	edges []Edge,
) ([][]int64, [][]bool, error) {
	dist := make([][]int64, n)
	reachable := make([][]bool, n)

	for i := 0; i < n; i++ {
		dist[i] = make([]int64, n)
		reachable[i] = make([]bool, n)
		reachable[i][i] = true // Distance to itself is zero.
	}

	// Initialize direct edges; keep the smallest parallel edge.
	for _, edge := range edges {
		u, v := edge.From, edge.To

		if !reachable[u][v] || edge.Weight < dist[u][v] {
			dist[u][v] = edge.Weight
			reachable[u][v] = true
		}
	}

	// k must be the outer loop.
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			if !reachable[i][k] {
				continue
			}

			for j := 0; j < n; j++ {
				if !reachable[k][j] {
					continue
				}

				candidate, ok := checkedAdd(dist[i][k], dist[k][j])
				if !ok {
					return nil, nil, fmt.Errorf("distance overflow")
				}

				if !reachable[i][j] || candidate < dist[i][j] {
					dist[i][j] = candidate
					reachable[i][j] = true
				}
			}
		}
	}

	// A negative distance back to itself implies a negative cycle.
	for i := 0; i < n; i++ {
		if dist[i][i] < 0 {
			return nil, nil, fmt.Errorf("negative cycle detected")
		}
	}

	return dist, reachable, nil
}
