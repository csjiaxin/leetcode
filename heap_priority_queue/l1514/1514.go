package l1514

import "leetcode/heap_priority_queue/internal/pqueue"

type edge struct {
	to          int
	probability float64
}
type state struct {
	node        int
	probability float64
}

// maxProbability uses a maximum heap; probabilities in [0,1] cannot improve along a path.
// O(n+E log(E+1)) time, O(n+E) space.
func maxProbability(n int, edges [][]int, succProb []float64, start_node, end_node int) float64 {
	graph := make([][]edge, n)
	for i, e := range edges {
		graph[e[0]] = append(graph[e[0]], edge{e[1], succProb[i]})
		graph[e[1]] = append(graph[e[1]], edge{e[0], succProb[i]})
	}
	best := make([]float64, n)
	best[start_node] = 1
	q := pqueue.New(func(a, b state) bool { return a.probability > b.probability }, state{start_node, 1})
	for q.Len() > 0 {
		s := q.Pop()
		if s.probability < best[s.node] {
			continue
		}
		if s.node == end_node {
			return s.probability
		}
		for _, e := range graph[s.node] {
			p := s.probability * e.probability
			if p > best[e.to] {
				best[e.to] = p
				q.Push(state{e.to, p})
			}
		}
	}
	return 0
}
