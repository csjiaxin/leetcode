// graph[u] contains the nodes that must come after u.
// Returns (order, true), or (nil, false) if a cycle exists.
func TopoSort(graph [][]int) ([]int, bool) {
	n := len(graph)
	indegree := make([]int, n)

	for _, neighbors := range graph {
		for _, v := range neighbors {
			indegree[v]++
		}
	}

	queue := make([]int, 0, n)
	for u := 0; u < n; u++ {
		if indegree[u] == 0 {
			queue = append(queue, u)
		}
	}

	order := make([]int, 0, n)

	for head := 0; head < len(queue); head++ {
		u := queue[head]
		order = append(order, u)

		for _, v := range graph[u] {
			indegree[v]--

			if indegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(order) != n {
		return nil, false // Some nodes are blocked by a cycle.
	}

	return order, true
}