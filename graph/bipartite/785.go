package main

func isBipartite(graph [][]int) bool {

	n := len(graph)
	// 0 unvisit, 1 party 1,-1 another party
	color := make([]int, n)

	for c := 0; c < n; c++ {
		if color[c] == 0 {
			color[c] = 1

			q := []int{c}
			for head := 0; head < len(q); head++ {
				visit_node := q[head]

				for _, neigh_node := range graph[visit_node] {
					if color[neigh_node] == 0 {
						color[neigh_node] = -color[visit_node]
						q = append(q, neigh_node)
					} else if color[neigh_node] == color[visit_node] {
						return false
					}

				}

			}
		}
	}
	return true
}
