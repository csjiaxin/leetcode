package main

import "fmt"

// https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/
func countSubTrees(n int, edges [][]int, labels string) []int {
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	answer := make([]int, n)
	if n == 0 {
		return answer
	}
	TreeFold(graph, 0, -1, func(node int) [26]int {
		var counts [26]int
		counts[labels[node]-'a'] = 1
		return counts
	}, func(left, right [26]int) [26]int {
		for i := range left {
			left[i] += right[i]
		}
		return left
	}, func(node int, counts [26]int) { answer[node] = counts[labels[node]-'a'] })
	return answer
}

// TreeFold 用于无向树；parent 防止沿父边回走，不适用于任意有环图。
func TreeFold[S any](graph [][]int, node, parent int, start func(int) S, merge func(S, S) S, finish func(int, S)) S {
	state := start(node)
	for _, child := range graph[node] {
		if child != parent {
			state = merge(state, TreeFold(graph, child, node, start, merge, finish))
		}
	}
	if finish != nil {
		finish(node, state)
	}
	return state
}

func main() {
	fmt.Println(countSubTrees(7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, "abaedcd"))
}
